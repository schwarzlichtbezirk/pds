package main

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	"github.com/schwarzlichtbezirk/pds/pb"
)

// ReadDataFile reads pds-ports.json file step by step,
// and sends readed ports to gRPC stream.
func ReadDataFile(fname string) (err error) {
	grpclog.Infof("read file '%s'\n", fname)

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
	if stream, err = grpcPort.RecordList(ctx); err != nil {
		return
	}

	// finally get stream result
	defer func() {
		var reply *pb.Summary
		if reply, err = stream.CloseAndRecv(); err != nil {
			return
		}
		grpclog.Infof("data base summary: readed %d ports, elapsed %dms",
			reply.PortCount, reply.ElapsedTime)
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
			return
		}
		if len(port.Coordinates) != 2 {
			grpclog.Warningf("port without coordinates: %s, %s\n", port.Unlocs[0], port.Name)
		}
	}

	// read closing bracket
	if _, err = dec.Token(); err != nil {
		return
	}

	return nil
}
