package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/jessevdk/go-flags"
	"github.com/schwarzlichtbezirk/pds/pb"
	"google.golang.org/grpc"
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

var (
	grpcClient pb.PortGuideClient
	grpcTool   pb.ToolGuideClient
)

// Init performs global data initialization.
func Init() {
	log.Println("starts")

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
				log.Println("shutting down by timeout")
			} else if errors.Is(exitctx.Err(), context.Canceled) {
				log.Println("shutting down by cancel")
			} else {
				log.Printf("shutting down by %s", exitctx.Err().Error())
			}
		case <-sigint:
			log.Println("shutting down by break")
		case <-sigterm:
			log.Println("shutting down by process termination")
		}
		signal.Stop(sigint)
		signal.Stop(sigterm)
	}()

	// load content of Config structure from YAML-file.
	if !cfg.NoConfig {
		var err error

		// get confiruration path
		if ConfigPath, err = DetectConfigPath(); err != nil {
			log.Fatal(err)
		}
		log.Printf("config path: %s\n", ConfigPath)

		if err = ReadYaml(cfgfile, &cfg); err != nil {
			log.Fatalf("can not read '%s' file: %v\n", cfgfile, err)
		}
		log.Printf("loaded '%s'\n", cfgfile)
		// second iteration, rewrite settings from config file
		if _, err = flags.NewParser(&cfg, flags.PassDoubleDash).Parse(); err != nil {
			panic("no way to here")
		}
	}
	// correct config
	cfg.AddrGRPC = envfmt(cfg.AddrGRPC)
}

// Run launches server listeners.
func Run(gmux *Router) {
	// helps to start HTTP only after gRPC to prevent call to uninitialized data
	var grpcctx, grpccancel = context.WithCancel(context.Background())
	var httpctx, httpcancel = context.WithCancel(context.Background())

	// starts gRPC client
	exitwg.Add(1)
	go func() {
		defer exitwg.Done()

		var err error
		var conn *grpc.ClientConn

		var addrs []resolver.Address
		for _, url := range strings.Split(envfmt(cfg.AddrGRPC), ";") {
			for _, port := range strings.Split(envfmt(cfg.PortGRPC), ";") {
				addrs = append(addrs, resolver.Address{Addr: url + port})
			}
		}
		var r = manual.NewBuilderWithScheme("pds")
		r.InitialState(resolver.State{
			Addresses: addrs,
		})

		const serviceConfig = `{"loadBalancingPolicy":"round_robin"}`
		var address = fmt.Sprintf("%s:///unused", r.Scheme())
		var options = []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithBlock(),
			grpc.WithResolvers(r),
			grpc.WithDefaultServiceConfig(serviceConfig),
		}

		log.Printf("grpc connecting on %s\n", address)
		go func() {
			defer grpccancel()
			if conn, err = grpc.DialContext(grpcctx, address, options...); err != nil {
				log.Fatalf("fail to dial on %s: %v", address, err)
			}
			grpcClient = pb.NewPortGuideClient(conn)
			grpcTool = pb.NewToolGuideClient(conn)

			log.Printf("grpc connected on %s\n", address)
		}()
		// wait until connect will be established or have got exit signal
		select {
		case <-grpcctx.Done(): // data is ready, so HTTP can safely serve
		case <-exitctx.Done():
			log.Printf("grpc connection canceled on %s\n", address)
			return
		}

		// wait for exit signal
		<-exitctx.Done()

		if err := conn.Close(); err != nil {
			log.Printf("grpc disconnect on %s: %v\n", address, err)
		} else {
			log.Printf("grpc disconnected on %s\n", address)
		}
	}()

	// starts HTTP listeners
	go func() {
		defer httpcancel()
		var httpwg sync.WaitGroup

		for _, addr := range cfg.PortHTTP {
			var addr = envfmt(addr) // localize
			httpwg.Add(1)
			exitwg.Add(1)
			go func() {
				defer exitwg.Done()

				var server = &http.Server{
					Addr:              addr,
					Handler:           gmux,
					ReadTimeout:       cfg.ReadTimeout,
					ReadHeaderTimeout: cfg.ReadHeaderTimeout,
					WriteTimeout:      cfg.WriteTimeout,
					IdleTimeout:       cfg.IdleTimeout,
					MaxHeaderBytes:    cfg.MaxHeaderBytes,
				}
				// wait until connect will be established or have got exit signal
				select {
				case <-grpcctx.Done():
				case <-exitctx.Done():
					return
				}

				log.Printf("start http on %s\n", addr)
				go func() {
					httpwg.Done()
					if err := server.ListenAndServe(); err != http.ErrServerClosed {
						log.Fatalf("failed to serve: %v", err)
					}
				}()

				// wait for exit signal
				<-exitctx.Done()

				// create a deadline to wait for.
				var ctx, cancel = context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
				defer cancel()

				server.SetKeepAlivesEnabled(false)
				if err := server.Shutdown(ctx); err != nil {
					log.Printf("shutdown http on %s: %v\n", addr, err)
				} else {
					log.Printf("stop http on %s\n", addr)
				}
			}()
		}

		httpwg.Wait()
	}()

	select {
	case <-grpcctx.Done():
		log.Printf("grpc ready")
	case <-exitctx.Done():
		return
	}

	select {
	case <-httpctx.Done():
		log.Printf("http ready")
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
	log.Println("shutting down complete.")
}
