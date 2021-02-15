package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	pb "github.com/schwarzlichtbezirk/pds-grpc/pds"
)

// Protobuf client error messages
var (
	ErrBadKey = errors.New("key token is not string")
)

// ReadDataFile reads ports.json file step by step,
// and sends readed ports to gRPC stream.
func ReadDataFile(fname string) (err error) {
	log.Printf("read file %s\n", fname)

	var f *os.File
	if f, err = os.Open(fname); err != nil {
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
		var t json.Token

		t, err = dec.Token()
		if _, ok := t.(string); !ok {
			return ErrBadKey
		}
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
