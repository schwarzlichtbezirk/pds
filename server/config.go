package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc/grpclog"
	"gopkg.in/yaml.v3"
)

// Instance of common service settings.
var cfg struct { // inits default values:
	ConfigPath string   `json:"-" yaml:"-" env:"CONFIGPATH" short:"c" long:"cfgpath" description:"Configuration path. Can be full path to config folder, or relative from executable destination."`
	NoConfig   bool     `json:"-" yaml:"-" long:"nocfg" description:"Specifies do not load settings from YAML-settings file, keeps default."`
	PortGRPC   []string `json:"port-grpc" yaml:"port-grpc" env:"PORTGRPC" env-delim:";" short:"p" long:"portgrpc" description:"List of ports of gRPC-services."`
}

func init() {
	if _, err := flags.Parse(&cfg); err != nil {
		os.Exit(1)
	}
}

const (
	gitname = "pds"
	gitpath = "github.com/schwarzlichtbezirk/" + gitname
	cfgbase = gitname + "-config"
	cfgfile = "server.yaml"
)

// ConfigPath determines configuration path, depended on what directory is exist.
var ConfigPath string

// ErrNoCongig is "no configuration path was found" error message.
var ErrNoCongig = errors.New("no configuration path was found")

// DetectConfigPath finds configuration path with existing configuration file at least.
func DetectConfigPath() (retpath string, err error) {
	var ok bool
	var path string
	var exepath = filepath.Dir(os.Args[0])

	// try to get from environment setting
	if cfg.ConfigPath != "" {
		path = envfmt(cfg.ConfigPath)
		// try to get access to full path
		if ok, _ = pathexists(filepath.Join(path, cfgfile)); ok {
			retpath = path
			return
		}
		// try to find relative from executable path
		path = filepath.Join(exepath, path)
		if ok, _ = pathexists(filepath.Join(path, cfgfile)); ok {
			retpath = path
			return
		}
		grpclog.Warningf("no access to pointed configuration path '%s'\n", cfg.ConfigPath)
	}

	// try to get from config subdirectory on executable path
	path = filepath.Join(exepath, cfgbase)
	if ok, _ = pathexists(filepath.Join(path, cfgfile)); ok {
		retpath = path
		return
	}
	// try to find in executable path
	if ok, _ = pathexists(filepath.Join(exepath, cfgfile)); ok {
		retpath = exepath
		return
	}
	// try to find in config subdirectory of current path
	if ok, _ = pathexists(filepath.Join(cfgbase, cfgfile)); ok {
		retpath = cfgbase
		return
	}
	// try to find in current path
	if ok, _ = pathexists(cfgfile); ok {
		retpath = "."
		return
	}
	// check up current path is the git root path
	if ok, _ = pathexists(filepath.Join("config", cfgfile)); ok {
		retpath = "config"
		return
	}

	// check up running in devcontainer workspace
	path = filepath.Join("/workspaces", gitname, "config")
	if ok, _ = pathexists(filepath.Join(path, cfgfile)); ok {
		retpath = path
		return
	}

	// check up git source path
	var prefix string
	if prefix, ok = os.LookupEnv("GOPATH"); ok {
		path = filepath.Join(prefix, "src", gitpath, "config")
		if ok, _ = pathexists(filepath.Join(path, cfgfile)); ok {
			retpath = path
			return
		}
	}

	// if GOBIN or GOPATH is present
	if prefix, ok = os.LookupEnv("GOBIN"); !ok {
		if prefix, ok = os.LookupEnv("GOPATH"); ok {
			prefix = filepath.Join(prefix, "bin")
		}
	}
	if ok {
		// try to get from go bin config
		path = filepath.Join(prefix, cfgbase)
		if ok, _ = pathexists(filepath.Join(path, cfgfile)); ok {
			retpath = path
			return
		}
		// try to get from go bin root
		if ok, _ = pathexists(filepath.Join(prefix, cfgfile)); ok {
			retpath = prefix
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
	if body, err = os.ReadFile(filepath.Join(ConfigPath, fname)); err != nil {
		return
	}
	if err = yaml.Unmarshal(body, data); err != nil {
		return
	}
	return
}
