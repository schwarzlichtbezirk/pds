package main

import (
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"gopkg.in/yaml.v3"
)

const (
	gitname = "pds"
	gitpath = "github.com/schwarzlichtbezirk/" + gitname
	cfgfile = "pds-server.yaml"
)

// CfgCmdLine is command line arguments managment settings.
type CfgCmdLine struct {
	ConfigPath string `json:"-" yaml:"-" env:"CONFIGPATH" short:"c" long:"cfgpath" description:"Configuration path. Can be full path to config folder, or relative from executable destination."`
	NoConfig   bool   `json:"-" yaml:"-" long:"nocfg" description:"Specifies do not load settings from YAML-settings file, keeps default."`
}

type CfgRpcServ struct {
	PortGRPC []string `json:"port-grpc" yaml:"port-grpc" env:"PORTGRPC" env-delim:";" short:"g" long:"portgrpc" description:"List of ports of gRPC-services."`
}

type CfgLogger struct {
	LogLevel        string `json:"log-level" yaml:"log-level" long:"ll" description:"The logging level the logger should log at. Can be: panic, fatal, error, warn, info, debug, trace."`
	ForceColors     bool   `json:"force-colors" yaml:"force-colors" long:"fc" description:"Set to true to bypass checking for a TTY before outputting colors."`
	TimestampFormat string `json:"timestamp-format,omitempty" yaml:"timestamp-format,omitempty" long:"tsf" description:"Format to use is the same than for time.Format or time.Parse from the standard library."`
}

// Config is common service settings.
type Config struct {
	CfgCmdLine `json:"-" yaml:"-" group:"Command line arguments"`
	CfgRpcServ `json:"grpc-server" yaml:"grpc-server" group:"gRPC Server"`
	CfgLogger  `json:"logger" yaml:"logger" group:"gRCP Logger"`
}

// Instance of common service settings.
var cfg = Config{ // inits default values:
	CfgRpcServ: CfgRpcServ{
		PortGRPC: []string{":50051", ":50052"},
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
