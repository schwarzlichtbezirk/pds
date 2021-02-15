package main

import (
	"context"
	"io"
	"log"
	"math"
	"strings"
	"sync"
	"time"

	pb "github.com/schwarzlichtbezirk/pds-grpc/pds"
)

// Haversine calculates distance in meters between two lati­tude/longi­tude points.
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371e3              // metres
	var φ1 = lat1 * math.Pi / 180 // φ, λ in radians
	var φ2 = lat2 * math.Pi / 180
	var Δφ = (lat2 - lat1) * math.Pi / 180
	var Δλ = (lon2 - lon1) * math.Pi / 180
	var a = math.Sin(Δφ/2)*math.Sin(Δφ/2) +
		math.Cos(φ1)*math.Cos(φ2)*
			math.Sin(Δλ/2)*math.Sin(Δλ/2)
	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = R * c // in metres
	return d
}

type routeGuideServer struct {
	pb.UnimplementedPortGuideServer
	addr string
}

// Storage is singleton, PDS database
var storage sync.Map

func (s *routeGuideServer) RecordList(stream pb.PortGuide_RecordListServer) error {
	var count int32
	var startTime = time.Now()
	for {
		var port, err = stream.Recv()
		if err == io.EOF {
			log.Printf("fetched %d items\n", count)
			var endTime = time.Now()
			return stream.SendAndClose(&pb.Summary{
				PortCount:   count,
				ElapsedTime: int32(endTime.Sub(startTime).Milliseconds()),
			})
		}
		if err != nil {
			return err
		}
		count++
		storage.Store(port.Unlocs[0], port)
	}
}

func (s *routeGuideServer) SetByKey(ctx context.Context, port *pb.Port) (*pb.Key, error) {
	var key = port.GetUnlocs()[0]
	storage.Store(key, port)
	return &pb.Key{Value: key}, nil
}

func (s *routeGuideServer) GetByKey(ctx context.Context, key *pb.Key) (*pb.Port, error) {
	if v, ok := storage.Load(key.Value); ok {
		return v.(*pb.Port), nil
	}
	return &pb.Port{}, nil
}

func (s *routeGuideServer) GetByName(ctx context.Context, name *pb.Name) (*pb.Port, error) {
	var found = &pb.Port{} // result
	storage.Range(func(key, val interface{}) bool {
		var port = val.(*pb.Port)
		if port.Name == name.Value {
			found = port
			return false
		}
		return true
	})
	return found, nil
}

func (s *routeGuideServer) FindNearest(ctx context.Context, coord *pb.Point) (*pb.Port, error) {
	var distance float64 = 1e10 // let's set it to any maximum possible value
	var found = &pb.Port{}      // result
	storage.Range(func(key, val interface{}) bool {
		var port = val.(*pb.Port)
		if len(port.Coordinates) == 2 {
			var d = Haversine(
				float64(coord.Latitude), float64(coord.Longitude),
				float64(port.Coordinates[1]), float64(port.Coordinates[0]))
			if d < distance {
				found, distance = port, d
			}
		}
		return true
	})
	return found, nil
}

func (s *routeGuideServer) FindInCircle(ctx context.Context, circ *pb.Circle) (*pb.Ports, error) {
	var ports = pb.Ports{}
	var r = float64(circ.Radius)
	storage.Range(func(key, val interface{}) bool {
		var port = val.(*pb.Port)
		if len(port.Coordinates) == 2 {
			var d = Haversine(
				float64(circ.Center.Latitude), float64(circ.Center.Longitude),
				float64(port.Coordinates[1]), float64(port.Coordinates[0]))
			if d < r {
				ports.List = append(ports.List, port)
			}
		}
		return true
	})
	return &ports, nil
}

func (s *routeGuideServer) FindText(ctx context.Context, q *pb.Quest) (*pb.Ports, error) {
	var sub = q.Value
	if !q.Sensitive {
		sub = strings.ToLower(sub)
	}
	var cmp = func(text string) bool {
		if !q.Sensitive {
			text = strings.ToLower(text)
		}
		if q.Whole {
			return text == sub
		}
		return strings.Contains(text, sub)
	}

	var ports = pb.Ports{}
	storage.Range(func(key, val interface{}) bool {
		var port = val.(*pb.Port)
		if cmp(port.Name) || cmp(port.City) || cmp(port.Province) || cmp(port.Country) {
			ports.List = append(ports.List, port)
		}
		return true
	})
	return &ports, nil
}