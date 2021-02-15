package main

// Config is common service settings.
type Config struct {
	GrpcPort []string `json:"grpc-port" yaml:"grpc-port"`
}

// Instance of common service settings.
var cfg = Config{ // inits default values:
	GrpcPort: []string{":50051", ":50052"},
}
