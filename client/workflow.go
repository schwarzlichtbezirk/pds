package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/resolver/manual"
)

var (
	// context to indicate about service shutdown
	exitctx context.Context
	exitfn  context.CancelFunc
	// wait group for all server goroutines
	exitwg sync.WaitGroup
)

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
}

// Init performs global data initialization.
func Init() {
	grpclog.Infoln("starts")

	// create context and wait the break
	exitctx, exitfn = context.WithCancel(context.Background())
	go func() {
		// Make exit signal on function exit.
		defer exitfn()

		var sigint = make(chan os.Signal, 1)
		var sigterm = make(chan os.Signal, 1)
		// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C) or SIGTERM (Ctrl+/)
		// SIGKILL, SIGQUIT will not be caught.
		signal.Notify(sigint, syscall.SIGINT)
		signal.Notify(sigterm, syscall.SIGTERM)
		// Block until we receive our signal.
		select {
		case <-exitctx.Done():
			if errors.Is(exitctx.Err(), context.DeadlineExceeded) {
				grpclog.Infoln("shutting down by timeout")
			} else if errors.Is(exitctx.Err(), context.Canceled) {
				grpclog.Infoln("shutting down by cancel")
			} else {
				grpclog.Infof("shutting down by %s", exitctx.Err().Error())
			}
		case <-sigint:
			grpclog.Infoln("shutting down by break")
		case <-sigterm:
			grpclog.Infoln("shutting down by process termination")
		}
		signal.Stop(sigint)
		signal.Stop(sigterm)
	}()

	// load content of Config structure from YAML-file.
	if !cfg.NoConfig {
		var err error

		// get confiruration path
		if ConfigPath, err = DetectConfigPath(); err != nil {
			grpclog.Fatal(err)
		}
		grpclog.Infof("config path: %s\n", ConfigPath)

		if err = ReadYaml(cfgfile, &cfg); err != nil {
			grpclog.Fatalf("can not read '%s' file: %v\n", cfgfile, err)
		}
		grpclog.Infof("loaded '%s'\n", cfgfile)
		// second iteration, rewrite settings from config file
		if _, err = flags.NewParser(&cfg, flags.PassDoubleDash).Parse(); err != nil {
			panic("no way to here")
		}
	}
	// correct config
	cfg.AddrGRPC = envfmt(cfg.AddrGRPC)
}

// Run launches server listeners.
func Run() {
	var grpcctx, grpccancel = context.WithCancel(context.Background())
	var mux = runtime.NewServeMux()

	// starts HTTP-gRPC proxy
	exitwg.Add(1)
	go func() {
		defer exitwg.Done()
		defer grpccancel() // send close signal to gRPC endpoint function

		var addrs []resolver.Address
		for _, addr := range strings.Split(envfmt(cfg.AddrGRPC), ";") {
			addrs = append(addrs, resolver.Address{Addr: addr})
		}
		var r = manual.NewBuilderWithScheme(cfg.SchemeGRPC)
		r.InitialState(resolver.State{
			Addresses: addrs,
		})

		const serviceConfig = `{"loadBalancingPolicy":"round_robin"}`
		var address = fmt.Sprintf("%s:///unused", r.Scheme())
		var options = []grpc.DialOption{
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
			grpc.WithResolvers(r),
			grpc.WithDefaultServiceConfig(serviceConfig),
		}

		// establish connection and create gRPC clients
		grpclog.Infof("grpc connecting on %s\n", address)
		if err := RegisterAllHandlersFromEndpoint(grpcctx, mux, address, options); err != nil {
			grpclog.Fatalf("failed to register gateway on %s: %v", address, err)
		}
		grpclog.Infof("grpc connected on %s\n", address)

		// init database
		if err := ReadDataFile(envfmt(cfg.DataFile)); err != nil {
			grpclog.Fatal(err)
		}

		// data is ready, so HTTP can safely serve
		var httpwg sync.WaitGroup
		for _, addr := range strings.Split(envfmt(cfg.PortHTTP), ";") {
			var addr = envfmt(addr) // localize
			httpwg.Add(1)
			exitwg.Add(1)
			go func() {
				defer exitwg.Done()

				var server = &http.Server{
					Addr:              addr,
					Handler:           mux,
					ReadTimeout:       cfg.ReadTimeout,
					ReadHeaderTimeout: cfg.ReadHeaderTimeout,
					WriteTimeout:      cfg.WriteTimeout,
					IdleTimeout:       cfg.IdleTimeout,
					MaxHeaderBytes:    cfg.MaxHeaderBytes,
				}

				grpclog.Infof("start http on %s\n", addr)
				go func() {
					httpwg.Done()
					if err := server.ListenAndServe(); err != http.ErrServerClosed {
						grpclog.Fatalf("failed to serve: %v", err)
					}
				}()

				// wait for exit signal
				<-exitctx.Done()

				// create a deadline to wait for.
				var ctx, cancel = context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
				defer cancel()

				server.SetKeepAlivesEnabled(false)
				if err := server.Shutdown(ctx); err != nil {
					grpclog.Errorf("shutdown http on %s: %v\n", addr, err)
				} else {
					grpclog.Infof("stop http on %s\n", addr)
				}
			}()
		}
		httpwg.Wait()
		grpclog.Infoln("ready")

		// wait for exit signal
		<-exitctx.Done()
		grpclog.Infoln("grpc disconnect")
	}()
}

// Done performs graceful network shutdown,
// waits until all server threads will be stopped.
func Done() {
	// wait for exit signal
	<-exitctx.Done()
	// wait until all server threads will be stopped.
	exitwg.Wait()
	// give opportunity to get and process close signal to all goroutines
	time.Sleep(5 * time.Millisecond)
	grpclog.Infoln("shutting down complete.")
}
