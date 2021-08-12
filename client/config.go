package main

import "time"

// CfgServ is web server settings.
type CfgServ struct {
	PortHTTP          []string      `json:"port-http" yaml:"port-http"`
	ReadTimeout       time.Duration `json:"read-timeout" yaml:"read-timeout"`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      time.Duration `json:"write-timeout" yaml:"write-timeout"`
	IdleTimeout       time.Duration `json:"idle-timeout" yaml:"idle-timeout"`
	MaxHeaderBytes    int           `json:"max-header-bytes" yaml:"max-header-bytes"`
	// Maximum duration to wait for graceful shutdown.
	ShutdownTimeout time.Duration `json:"shutdown-timeout" yaml:"shutdown-timeout"`
}

// Config is common service settings.
type Config struct {
	CfgServ  `json:"webserver" yaml:"webserver"`
	AddrGRPC string `json:"addr-grpc" yaml:"addr-grpc"`
	PortGRPC string `json:"port-grpc" yaml:"port-grpc"`
	DataFile string `json:"data-file" yaml:"data-file"`
}

// Instance of common service settings.
var cfg = Config{ // inits default values:
	CfgServ: CfgServ{
		PortHTTP:          []string{":8008"},
		ReadTimeout:       time.Duration(15) * time.Second,
		ReadHeaderTimeout: time.Duration(15) * time.Second,
		WriteTimeout:      time.Duration(15) * time.Second,
		IdleTimeout:       time.Duration(60) * time.Second,
		MaxHeaderBytes:    1 << 20,
		ShutdownTimeout:   time.Duration(15) * time.Second,
	},
	AddrGRPC: "${PDSSERVURL}",
	PortGRPC: ":50051;:50052",
	DataFile: "ports.json",
}
