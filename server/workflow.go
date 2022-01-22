package main

import (
	"context"
	"errors"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/schwarzlichtbezirk/pds/pb"

	grcplogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
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
	grcplogrus.ReplaceGrpcLogger(grpclog)
}

// Init performs global data initialization.
func Init() {
	grpclog.Printf("version: %s, builton: %s\n", buildvers, builddate)
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
		// second logger setup - with updated config values
		SetupLogger()
	}
}

// Run launches server listeners.
func Run() {
	// starts gRPC servers
	var grpcctx, grpccancel = context.WithCancel(context.Background())
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
		grpccancel()
	}()

	// wait until exit or service is ready
	select {
	case <-grpcctx.Done():
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
	grpclog.Infoln("shutting down complete.")
}
