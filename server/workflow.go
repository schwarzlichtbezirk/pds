package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/schwarzlichtbezirk/pds/pb"
	"google.golang.org/grpc"
)

var (
	// context to indicate about service shutdown
	exitctx context.Context
	exitfn  context.CancelFunc
	// wait group for all server goroutines
	exitwg sync.WaitGroup
)

// Init performs global data initialization.
func Init() {
	log.Println("starts")

	flag.Parse()

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

	var err error

	// get confiruration path
	if ConfigPath, err = DetectConfigPath(); err != nil {
		log.Fatal(err)
	}
	log.Printf("config path: %s\n", ConfigPath)

	// load content of Config structure from YAML-file.
	if err = ReadYaml(cfgfile, &cfg); err != nil {
		log.Fatalf("can not read '%s' file: %v\n", cfgfile, err)
	}
	log.Printf("loaded '%s'\n", cfgfile)
}

// Run launches server listeners.
func Run() {
	// starts gRPC servers
	for _, addr := range cfg.PortGRPC {
		var addr = addr // localize
		exitwg.Add(1)
		go func() {
			defer exitwg.Done()

			var err error
			var lis net.Listener

			log.Printf("grpc server %s starts\n", addr)
			if lis, err = net.Listen("tcp", addr); err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			var server = grpc.NewServer()
			pb.RegisterToolGuideServer(server, &routeToolGuideServer{addr: addr})
			pb.RegisterPortGuideServer(server, &routePortGuideServer{addr: addr})
			go func() {
				if err = server.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
			}()

			// wait for exit signal
			<-exitctx.Done()

			server.GracefulStop()

			log.Printf("grpc server %s closed\n", addr)
		}()
	}

	log.Println("ready")
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
