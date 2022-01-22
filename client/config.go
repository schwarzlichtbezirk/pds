package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v3"
)

const (
	gitname = "pds"
	gitpath = "github.com/schwarzlichtbezirk/" + gitname
	cfgfile = "pds-client.yaml"
)

// CfgCmdLine is command line arguments managment settings.
type CfgCmdLine struct {
	ConfigPath string `json:"-" yaml:"-" env:"CONFIGPATH" short:"c" long:"cfgpath" description:"Configuration path. Can be full path to config folder, or relative from executable destination."`
	NoConfig   bool   `json:"-" yaml:"-" long:"nocfg" description:"Specifies do not load settings from YAML-settings file, keeps default."`
}

// CfgData is data managment settings.
type CfgDataKit struct {
	DataFile string `json:"data-file" yaml:"data-file" short:"d" long:"data" description:"Name of file with database."`
}

// CfgWebServ is web server settings.
type CfgWebServ struct {
	PortHTTP          []string      `json:"port-http" yaml:"port-http" env:"PORTHTTP" env-delim:";" short:"w" long:"http" description:"List of address:port values for non-encrypted connections. Address is skipped in most common cases, port only remains."`
	ReadTimeout       time.Duration `json:"read-timeout" yaml:"read-timeout" long:"rt" description:"Maximum duration for reading the entire request, including the body."`
	ReadHeaderTimeout time.Duration `json:"read-header-timeout" yaml:"read-header-timeout" long:"rht" description:"Amount of time allowed to read request headers."`
	WriteTimeout      time.Duration `json:"write-timeout" yaml:"write-timeout" long:"wt" description:"Maximum duration before timing out writes of the response."`
	IdleTimeout       time.Duration `json:"idle-timeout" yaml:"idle-timeout" long:"it" description:"Maximum amount of time to wait for the next request when keep-alives are enabled."`
	MaxHeaderBytes    int           `json:"max-header-bytes" yaml:"max-header-bytes" long:"mhb" description:"Controls the maximum number of bytes the server will read parsing the request header's keys and values, including the request line, in bytes."`
	// Maximum duration to wait for graceful shutdown.
	ShutdownTimeout time.Duration `json:"shutdown-timeout" yaml:"shutdown-timeout" long:"st" description:"Maximum duration to wait for graceful shutdown."`
}

type CfgRpcServ struct {
	AddrGRPC   []string `json:"addr-grpc" yaml:"addr-grpc" env:"ADDRGRPC" env-delim:";" short:"g" long:"grcp" description:"List of URL or IP-addresses with gRPC-services hosts."`
	SchemeGRPC string   `json:"scheme-grpc,omitempty" yaml:"scheme-grpc,omitempty" long:"scheme" description:"gRPC scheme name."`
}

type CfgLogger struct {
	LogLevel        string `json:"log-level" yaml:"log-level" long:"ll" description:"The logging level the logger should log at. Can be: panic, fatal, error, warn, info, debug, trace."`
	ForceColors     bool   `json:"force-colors" yaml:"force-colors" long:"fc" description:"Set to true to bypass checking for a TTY before outputting colors."`
	TimestampFormat string `json:"timestamp-format,omitempty" yaml:"timestamp-format,omitempty" long:"tsf" description:"Format to use is the same than for time.Format or time.Parse from the standard library."`
}

// Config is common service settings.
type Config struct {
	CfgCmdLine `json:"-" yaml:"-" group:"Command line arguments"`
	CfgDataKit `json:"data-kit" yaml:"data-kit" group:"Data Parameters"`
	CfgWebServ `json:"web-server" yaml:"web-server" group:"Web Server"`
	CfgRpcServ `json:"grpc-server" yaml:"grpc-server" group:"gRPC Server"`
	CfgLogger  `json:"logger" yaml:"logger" group:"gRCP Logger"`
}

// Instance of common service settings.
var cfg = Config{ // inits default values:
	CfgDataKit: CfgDataKit{
		DataFile: "pds-ports.json",
	},
	CfgWebServ: CfgWebServ{
		PortHTTP:          []string{":8008"},
		ReadTimeout:       time.Duration(15) * time.Second,
		ReadHeaderTimeout: time.Duration(15) * time.Second,
		WriteTimeout:      time.Duration(15) * time.Second,
		IdleTimeout:       time.Duration(60) * time.Second,
		MaxHeaderBytes:    1 << 20,
		ShutdownTimeout:   time.Duration(15) * time.Second,
	},
	CfgRpcServ: CfgRpcServ{
		AddrGRPC:   []string{"localhost:50051", "localhost:50052"},
		SchemeGRPC: "pds",
	},
	CfgLogger: CfgLogger{
		LogLevel:        "info",
		ForceColors:     true,
		TimestampFormat: "15:04:05",
	},
}

// compiled binary version, sets by compiler with command
//    go build -ldflags="-X 'main.buildvers=%buildvers%'"
var buildvers string

// compiled binary build date, sets by compiler with command
//    go build -ldflags="-X 'main.builddate=%date%'"
var builddate string

func init() {
	if _, err := flags.Parse(&cfg); err != nil {
		os.Exit(1)
	}
}

// ReadYaml reads "data" object from YAML-file with given file name.
func ReadYaml(fname string, data interface{}) (err error) {
	var body []byte
	if body, err = os.ReadFile(filepath.Join(ConfigPath, fname)); err != nil {
		return
	}
	if err = yaml.Unmarshal(body, data); err != nil {
		return
	}
	return
}
