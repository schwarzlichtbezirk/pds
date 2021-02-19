package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	pb "github.com/schwarzlichtbezirk/pds-grpc/pds"
	"google.golang.org/grpc"
)

var (
	// channel to indicate about server shutdown
	exitchan chan struct{}
	// wait group for all server goroutines
	exitwg sync.WaitGroup
)

var (
	grpcClient pb.PortGuideClient
	grpcTool   pb.ToolGuideClient
)

// Run launches server listeners.
func Run(gmux *Router) {
	makeServerLabel("gRPC-PDS", "0.1.0")

	// check up PDSBACKURL environment variable
	if os.Getenv("PDSBACKURL") == "" {
		os.Setenv("PDSBACKURL", "localhost")
	}

	// inits exit channel
	exitchan = make(chan struct{})

	// helps to start HTTP only after gRPC to prevent call to uninitialized data
	var grpcready = make(chan struct{})

	// starts HTTP servers
	for _, addr := range cfg.PortHTTP {
		var addr = envfmt(addr) // localize
		exitwg.Add(1)
		go func() {
			defer exitwg.Done()

			var server = &http.Server{
				Addr:              addr,
				Handler:           gmux,
				ReadTimeout:       time.Duration(cfg.ReadTimeout) * time.Second,
				ReadHeaderTimeout: time.Duration(cfg.ReadHeaderTimeout) * time.Second,
				WriteTimeout:      time.Duration(cfg.WriteTimeout) * time.Second,
				IdleTimeout:       time.Duration(cfg.IdleTimeout) * time.Second,
				MaxHeaderBytes:    cfg.MaxHeaderBytes,
			}
			go func() {
				// wait until database will be initialized, and start to receive connections
				// or until exit is signaled
				select {
				case <-grpcready:
				case <-exitchan:
					return
				}
				log.Printf("web server %s starts\n", addr)
				if err := server.ListenAndServe(); err != http.ErrServerClosed {
					log.Fatalf("failed to serve: %v", err)
				}
			}()

			// wait for exit signal
			<-exitchan

			// create a deadline to wait for.
			var ctx, cancel = context.WithTimeout(
				context.Background(),
				time.Duration(cfg.ShutdownTimeout)*time.Second)
			defer cancel()

			server.SetKeepAlivesEnabled(false)
			if err := server.Shutdown(ctx); err != nil {
				log.Printf("HTTP server shutdown: %v\n", err)
			} else {
				log.Printf("web server %s closed\n", addr)
			}
		}()
	}

	// starts gRPC client
	exitwg.Add(1)
	go func() {
		defer exitwg.Done()

		var err error
		var conn *grpc.ClientConn

		var addr = envfmt(cfg.AddrGrpc)
		log.Printf("grpc connecting to %s\n", addr)
		if conn, err = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock()); err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		grpcClient = pb.NewPortGuideClient(conn)
		grpcTool = pb.NewToolGuideClient(conn)

		log.Printf("grpc connected to %s\n", addr)

		if err := ReadDataFile(envfmt(cfg.DataFile)); err != nil {
			log.Fatal(err)
		}

		// data is ready, so HTTP can safely serve
		close(grpcready)

		// wait for exit signal
		<-exitchan

		if err := conn.Close(); err != nil {
			log.Printf("gRPC disconnect: %v\n", err)
		} else {
			log.Println("grpc disconnected")
		}
	}()
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
