package main

// Config is common service settings.
type Config struct {
	PortGRPC []string `json:"port-grpc" yaml:"port-grpc"`
}

// Instance of common service settings.
var cfg = Config{ // inits default values:
	PortGRPC: []string{":50051", ":50052"},
}
