package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

var grpclog *logrus.Entry

func init() {
	// first setup logger with default config values
	SetupLogger()
}

// SetupLogger inits logger configured with service settings.
func SetupLogger() {
	var ll, err = logrus.ParseLevel(cfg.LogLevel)
	if err != nil {
		ll = logrus.InfoLevel
	}
	grpclog = logrus.NewEntry(&logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			ForceColors:     cfg.ForceColors,
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: cfg.TimestampFormat,
		},
		Hooks: make(logrus.LevelHooks),
		Level: ll,
	})
	grpc_logrus.ReplaceGrpcLogger(grpclog)
}

// Init performs global data initialization.
func Init() {
	grpclog.Infof("version: %s, builton: %s\n", buildvers, builddate)
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
		// second logger setup - with updated config values
		SetupLogger()
	}
}

// Run launches server listeners.
func Run() {
	var grpcctx, grpccancel = context.WithCancel(context.Background())
	var httpctx, httpcancel = context.WithCancel(context.Background())
	var mux = runtime.NewServeMux()

	// starts HTTP-gRPC proxy
	exitwg.Add(1)
	go func() {
		defer exitwg.Done()
		defer grpccancel() // send close signal to gRPC endpoint function

		var addrs []resolver.Address
		for _, addr := range cfg.AddrGRPC {
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
			//grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(monitoringClientUnary, retryUnary)),
			//grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(monitoringClientStream, retryStream)),
		}

		// establish connection and create gRPC clients
		grpclog.Infof("grpc connecting on %s\n", address)
		if err := RegisterAllHandlersFromEndpoint(grpcctx, mux, address, options); err != nil {
			grpclog.Fatalf("failed to register gateway on %s: %v", address, err)
		}
		grpclog.Infof("grpc connected on %s\n", address)

		// init database
		if err := ReadDataFile(EnvFmt(cfg.DataFile)); err != nil {
			grpclog.Fatal(err)
		}

		// data is ready, so HTTP can safely serve
		var httpwg sync.WaitGroup
		for _, addr := range cfg.PortHTTP {
			var addr = EnvFmt(addr) // localize
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
		httpcancel()

		// wait for exit signal
		<-exitctx.Done()
		grpclog.Infoln("grpc disconnect")
	}()

	// wait until exit or service is ready
	select {
	case <-httpctx.Done():
		grpclog.Infoln("service ready")
	case <-exitctx.Done():
		return
	}
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
