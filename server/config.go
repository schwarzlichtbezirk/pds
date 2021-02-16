package main

// Config is common service settings.
type Config struct {
	PortGrpc []string `json:"port-grpc" yaml:"port-grpc"`
}

// Instance of common service settings.
var cfg = Config{ // inits default values:
	PortGrpc: []string{":50051", ":50052"},
}
