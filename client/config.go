package main

// CfgServ is web server settings.
type CfgServ struct {
	PortHTTP          []string `json:"port-http" yaml:"port-http"`
	ReadTimeout       int      `json:"read-timeout" yaml:"read-timeout"`
	ReadHeaderTimeout int      `json:"read-header-timeout" yaml:"read-header-timeout"`
	WriteTimeout      int      `json:"write-timeout" yaml:"write-timeout"`
	IdleTimeout       int      `json:"idle-timeout" yaml:"idle-timeout"`
	MaxHeaderBytes    int      `json:"max-header-bytes" yaml:"max-header-bytes"`
	ShutdownTimeout   int      `json:"shutdown-timeout" yaml:"shutdown-timeout"`
}

// Config is common service settings.
type Config struct {
	CfgServ  `json:"webserver" yaml:"webserver"`
	AddrGrpc string `json:"addr-grpc" yaml:"addr-grpc"`
	DataFile string `json:"data-file" yaml:"data-file"`
}

// Instance of common service settings.
var cfg = Config{ // inits default values:
	CfgServ: CfgServ{
		PortHTTP:          []string{":8008"},
		ReadTimeout:       15,
		ReadHeaderTimeout: 15,
		WriteTimeout:      15,
		IdleTimeout:       60,
		MaxHeaderBytes:    1 << 20,
		ShutdownTimeout:   15,
	},
	AddrGrpc: "${SERVERHOST}:50052",
	DataFile: "${GOPATH}/src/github.com/schwarzlichtbezirk/pds-grpc/config/ports.json",
}
