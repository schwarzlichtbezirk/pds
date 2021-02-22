package main

import (
	"context"
	"testing"
	"time"

	pb "github.com/schwarzlichtbezirk/pds-grpc/pds"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/proto"
)

// Initial sample data to setup on server.
// There is 3 ports grouped around Dubai and 1 port in Miami.
var origPort = []*pb.Port{
	{
		Name: "Dubai",
		Coordinates: []float32{
			55.27,
			25.25,
		},
		City:     "Dubai",
		Province: "Dubayy [Dubai]",
		Country:  "United Arab Emirates",
		Alias:    []string{},
		Regions:  []string{},
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEDXB",
		},
		Code: "52005",
	},
	{
		Name:    "Port Rashid",
		City:    "Port Rashid",
		Country: "United Arab Emirates",
		Alias:   []string{},
		Regions: []string{},
		Coordinates: []float32{
			55.2756505,
			25.284755,
		},
		Province: "Dubai",
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AEPRA",
		},
		Code: "52005",
	},
	{
		Name: "Sharjah",
		Coordinates: []float32{
			55.38,
			25.35,
		},
		City:     "Sharjah",
		Province: "Ash Shariqah [Sharjah]",
		Country:  "United Arab Emirates",
		Alias:    []string{},
		Regions:  []string{},
		Timezone: "Asia/Dubai",
		Unlocs: []string{
			"AESHJ",
		},
		Code: "52070",
	},
	{
		Name:     "Miami",
		City:     "Miami",
		Province: "Florida",
		Country:  "United States",
		Alias:    []string{},
		Regions:  []string{},
		Coordinates: []float32{
			-80.1917902,
			25.7616798,
		},
		Timezone: "America/New_York",
		Unlocs: []string{
			"USMIA",
		},
		Code: "5201",
	},
}
var dubai = origPort[0]

func Transactions(t *testing.T) {
	var (
		err        error
		grpcConn   *grpc.ClientConn
		grpcClient pb.PortGuideClient
	)

	// make client connection for gRPC on localhost
	if len(cfg.PortGRPC) == 0 {
		t.Fatal("no any grpc port defined")
	}
	if grpcConn, err = grpc.Dial("localhost"+cfg.PortGRPC[0], grpc.WithInsecure(), grpc.WithBlock()); err != nil {
		t.Fatalf("fail to dial: %v", err)
	}
	defer func() {
		if err = grpcConn.Close(); err != nil {
			t.Fatalf("fail on disconnect: %v", err)
		}
		t.Logf("client disconnected")
	}()
	grpcClient = pb.NewPortGuideClient(grpcConn)
	t.Logf("client connected")

	// limit execution time of the gRPC calls
	var ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// test api core for /api/port/set
	var port *pb.Port
	for _, port = range origPort {
		var key *pb.Key
		if key, err = grpcClient.SetByKey(ctx, port); err != nil {
			t.Fatalf("fail on SetByKey for '%s' call: %v", port.Name, err)
		}
		if key.Value != port.Unlocs[0] {
			t.Errorf("returned key is not equal to original for '%s'", port.Name)
		}
	}

	// test api core for /api/port/get
	if port, err = grpcClient.GetByKey(ctx, &pb.Key{Value: "AEDXB"}); err != nil {
		t.Fatalf("fail on GetByKey call: %v", err)
	}
	if !proto.Equal(port, dubai) {
		t.Error("received by GetByKey object is not expected Dubai port")
	}

	// test api core for /api/port/name
	if port, err = grpcClient.GetByName(ctx, &pb.Name{Value: "Dubai"}); err != nil {
		t.Fatalf("fail on GetByName call: %v", err)
	}
	if !proto.Equal(port, dubai) {
		t.Error("received by GetByName object is not expected Dubai port")
	}

	// test api core for /api/port/near
	var p = pb.Point{
		Latitude:  25.229789,
		Longitude: 55.165100,
	}
	if port, err = grpcClient.FindNearest(ctx, &p); err != nil {
		t.Fatalf("fail on FindNearest call: %v", err)
	}
	if !proto.Equal(port, dubai) {
		t.Error("received by FindNearest object is not expected Dubai port")
	}

	// test api core for /api/port/circle
	var circ = pb.Circle{
		Center: &pb.Point{
			Latitude:  25.458155,
			Longitude: 55.148621,
		},
		Radius: 40000,
	}
	var ports *pb.Ports
	if ports, err = grpcClient.FindInCircle(ctx, &circ); err != nil {
		t.Fatalf("fail on FindInCircle call: %v", err)
	}
	if len(ports.List) != 3 {
		t.Errorf("FindInCircle should find 3 ports, found %d", len(ports.List))
	}
	for _, port = range ports.List {
		if port.Name == "Miami" {
			t.Error("Miami should not be found, it outside of circle")
		}
	}

	// test api core for /api/port/text #1
	var q1 = pb.Quest{
		Value:     "dubai",
		Sensitive: false,
	}
	if ports, err = grpcClient.FindText(ctx, &q1); err != nil {
		t.Fatalf("fail on FindText call: %v", err)
	}
	if len(ports.List) != 2 {
		t.Errorf("FindText should find 2 ports for 'dubai' search, found %d", len(ports.List))
	}
	for _, port = range ports.List {
		if port.Name == "Miami" {
			t.Error("Miami should not be found for 'dubai' search")
		}
		if port.Name == "Sharjah" {
			t.Error("Sharjah should not be found for 'dubai' search")
		}
	}

	// test api core for /api/port/text #2
	var q2 = pb.Quest{
		Value:     "flor",
		Sensitive: false,
		Whole:     false,
	}
	if ports, err = grpcClient.FindText(ctx, &q2); err != nil {
		t.Fatalf("fail on FindText call: %v", err)
	}
	if len(ports.List) != 1 {
		t.Errorf("FindText should find 1 ports for 'flor' search, found %d", len(ports.List))
	}
	if len(ports.List) > 0 && ports.List[0].Name != "Miami" {
		t.Error("Miami port not found for 'flor' search")
	}
}

func TestGRPC(t *testing.T) {
	t.Logf("starts")
	Run()
	t.Logf("run transactions")
	Transactions(t)
	t.Logf("shutting down begin")
	Shutdown()
	t.Logf("shutting down complete.")
}
