package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v3"
)

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
	AddrGRPC: "${SERVERURL}",
	PortGRPC: ":50051;:50052",
	DataFile: "ports.json",
}

const (
	cfgenv  = "CONFIGPATH"
	cfgfile = "client.yaml"
	cfgbase = "pds-config"
	srcpath = "src/github.com/schwarzlichtbezirk/pds/config"
)

// Path given from command line parameter.
var cmdpath = flag.String("d", "", "configuration path")

// ConfigPath determines configuration path, depended on what directory is exist.
var ConfigPath string

// ErrNoCongig is "no configuration path was found" error message.
var ErrNoCongig = errors.New("no configuration path was found")

// DetectConfigPath finds configuration path with existing configuration file at least.
func DetectConfigPath() (cfgpath string, err error) {
	var path string
	var exepath = filepath.Dir(os.Args[0])

	// try to get from environment setting
	if path = envfmt(os.Getenv(cfgenv)); path != "" {
		// try to get access to full path
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			cfgpath = path
			return
		}
		// try to find relative from executable path
		path = filepath.Join(exepath, path)
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			cfgpath = exepath
			return
		}
		log.Printf("no access to pointed configuration path '%s'\n", path)
	}

	// try to get from command path arguments
	if path = *cmdpath; path != "" {
		// try to get access to full path
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			cfgpath = path
			return
		}
		// try to find relative from executable path
		path = filepath.Join(exepath, path)
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			cfgpath = exepath
			return
		}
	}

	// try to get from config subdirectory on executable path
	path = filepath.Join(exepath, cfgbase)
	if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
		cfgpath = path
		return
	}
	// try to find in executable path
	if ok, _ := pathexists(filepath.Join(exepath, cfgfile)); ok {
		cfgpath = exepath
		return
	}
	// try to find in current path
	if ok, _ := pathexists(cfgfile); ok {
		cfgpath = "."
		return
	}

	// if GOPATH is present
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		// try to get from go bin config
		path = filepath.Join(gopath, "bin", cfgbase)
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			cfgpath = path
			return
		}
		// try to get from go bin root
		path = filepath.Join(gopath, "bin")
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			cfgpath = path
			return
		}
		// try to get from source code
		path = filepath.Join(gopath, srcpath)
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			cfgpath = path
			return
		}
	}

	// no config was found
	err = ErrNoCongig
	return
}

// ReadYaml reads "data" object from YAML-file with given file path.
func ReadYaml(fname string, data interface{}) (err error) {
	var body []byte
	if body, err = ioutil.ReadFile(filepath.Join(ConfigPath, fname)); err != nil {
		return
	}
	if err = yaml.Unmarshal(body, data); err != nil {
		return
	}
	return
}
