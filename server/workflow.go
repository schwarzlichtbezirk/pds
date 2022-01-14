package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jessevdk/go-flags"
	"github.com/schwarzlichtbezirk/pds/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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
				grpclog.Infof("shutting down by %s\n", exitctx.Err().Error())
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
}

// Run launches server listeners.
func Run() {
	// starts gRPC servers
	func() {
		var grpcwg sync.WaitGroup
		for _, addr := range cfg.PortGRPC {
			var addr = addr // localize
			grpcwg.Add(1)
			exitwg.Add(1)
			go func() {
				defer exitwg.Done()

				var err error
				var lis net.Listener

				grpclog.Infof("grpc server %s starts\n", addr)
				if lis, err = net.Listen("tcp", addr); err != nil {
					grpclog.Fatalf("failed to listen: %v", err)
				}
				var server = grpc.NewServer()
				pb.RegisterToolGuideServer(server, &routeToolGuideServer{addr: addr})
				pb.RegisterPortGuideServer(server, &routePortGuideServer{addr: addr})
				go func() {
					grpcwg.Done()
					if err = server.Serve(lis); err != nil {
						grpclog.Fatalf("failed to serve: %v", err)
					}
				}()

				// wait for exit signal
				<-exitctx.Done()

				server.GracefulStop()

				grpclog.Infof("grpc server %s closed\n", addr)
			}()
		}

		grpcwg.Wait()
		grpclog.Infoln("grpc ready")
	}()
}

// Done performs graceful network shutdown,
// waits until all server threads will be stopped.
func Done() {
	// wait for exit signal
	<-exitctx.Done()
	// wait until all server threads will be stopped.
	exitwg.Wait()
	grpclog.Infoln("shutting down complete.")
}
