package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"path/filepath"
	"time"

	pb "github.com/schwarzlichtbezirk/pds-grpc/pds"
)

// ConfigPath determines configuration path, depended on what directory is exist.
var ConfigPath string

// DetectConfigPath finds configuration path.
func DetectConfigPath() {
	var path string
	// try to get from environment setting
	if path = os.Getenv("APPCONFIGPATH"); path != "" {
		if ok, _ := pathexists(path); ok {
			ConfigPath = path
			return
		}
		log.Printf("no access to pointed configuration path '%s'\n", path)
	}
	// try to get from config subdirectory on executable path
	var exepath = filepath.Dir(os.Args[0])
	path = filepath.Join(exepath, "config")
	if ok, _ := pathexists(filepath.Join(path, cfg.DataFile)); ok {
		ConfigPath = path
		return
	}
	// try to find in executable path
	if ok, _ := pathexists(filepath.Join(exepath, cfg.DataFile)); ok {
		ConfigPath = exepath
		return
	}
	// try to find in current path
	if ok, _ := pathexists(cfg.DataFile); ok {
		ConfigPath = "."
		return
	}

	// if GOPATH is present
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		// try to get from go bin config
		path = filepath.Join(gopath, "bin/config")
		if ok, _ := pathexists(filepath.Join(path, cfg.DataFile)); ok {
			ConfigPath = path
			return
		}
		// try to get from go bin root
		path = filepath.Join(gopath, "bin")
		if ok, _ := pathexists(filepath.Join(path, cfg.DataFile)); ok {
			ConfigPath = path
			return
		}
		// try to get from source code
		path = filepath.Join(gopath, "src/github.com/schwarzlichtbezirk/pds-grpc/config")
		if ok, _ := pathexists(path); ok {
			ConfigPath = path
			return
		}
	}

	// no config was found
	log.Fatal("no configuration path was found")
}

// ReadDataFile reads ports.json file step by step,
// and sends readed ports to gRPC stream.
func ReadDataFile(fname string) (err error) {
	log.Printf("read file %s\n", fname)

	var f *os.File
	if f, err = os.Open(filepath.Join(ConfigPath, fname)); err != nil {
		return
	}
	defer f.Close()

	// limit execution time of the action
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// inits gRPC stream
	var stream pb.PortGuide_RecordListClient
	if stream, err = grpcClient.RecordList(ctx); err != nil {
		log.Fatalf("%v.RecordList(_) = _, %v", grpcClient, err)
	}

	// finally get stream result
	defer func() {
		var reply *pb.Summary
		if reply, err = stream.CloseAndRecv(); err != nil {
			log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
		}
		log.Printf("data file summary: %v", reply)
	}()

	var dec = json.NewDecoder(f)

	// read open bracket
	if _, err = dec.Token(); err != nil {
		return
	}

	// while the array contains values
	for dec.More() {
		var port pb.Port

		// read and skip key token
		if _, err = dec.Token(); err != nil {
			return
		}
		// read port structure
		if err = dec.Decode(&port); err != nil {
			return
		}
		if err = stream.Send(&port); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, &port, err)
		}
		if len(port.Coordinates) != 2 {
			log.Printf("port without coordinates: %s, %s\n", port.Unlocs[0], port.Name)
		}
	}

	// read closing bracket
	if _, err = dec.Token(); err != nil {
		return
	}

	return nil
}
