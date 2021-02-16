package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	pb "github.com/schwarzlichtbezirk/pds-grpc/pds"
	"google.golang.org/grpc"
)

var (
	// channel to indicate about server shutdown
	exitchan chan struct{}
	// wait group for all server goroutines
	exitwg sync.WaitGroup
)

// Run launches server listeners.
func Run() {
	// inits exit channel
	exitchan = make(chan struct{})

	// starts gRPC servers
	for _, addr := range cfg.PortGrpc {
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
			pb.RegisterPortGuideServer(server, &routeGuideServer{addr: addr})
			go func() {
				if err = server.Serve(lis); err != nil {
					log.Fatalf("failed to serve: %v", err)
				}
			}()

			// wait for exit signal
			<-exitchan

			server.GracefulStop()

			log.Printf("grpc server %s closed\n", addr)
		}()
	}
}

// WaitBreak blocks goroutine until Ctrl+C will be pressed.
func WaitBreak() {
	var sigint = make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C) or SIGTERM (Ctrl+/)
	// SIGKILL, SIGQUIT will not be caught.
	signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-sigint
}

// Shutdown performs graceful network shutdown,
// waits until all server threads will be stopped.
func Shutdown() {
	close(exitchan)
	exitwg.Wait()
}
