package main

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const cfgfile = "server.yaml"

// ConfigPath determines configuration path, depended on what directory is exist.
var ConfigPath string

func pathexists(path string) (bool, error) {
	var err error
	if _, err = os.Stat(path); err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// DetectConfigPath finds configuration path.
func DetectConfigPath() {
	var path string
	// try to get from environment setting
	if path = os.Getenv("APPCONFIGPATH"); path != "" {
		if ok, _ := pathexists(path); ok {
			ConfigPath = path
			return
		}
		log.Printf("no access to pointed configuration path '%s'\n", path)
	}
	// try to get from config subdirectory on executable path
	var exepath = filepath.Dir(os.Args[0])
	path = filepath.Join(exepath, "pds-config")
	if ok, _ := pathexists(path); ok {
		ConfigPath = path
		return
	}
	// try to find in executable path
	if ok, _ := pathexists(filepath.Join(exepath, cfgfile)); ok {
		ConfigPath = exepath
		return
	}
	// try to find in current path
	if ok, _ := pathexists(cfgfile); ok {
		ConfigPath = "."
		return
	}

	// if GOPATH is present
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		// try to get from go bin config
		path = filepath.Join(gopath, "bin/pds-config")
		if ok, _ := pathexists(path); ok {
			ConfigPath = path
			return
		}
		// try to get from go bin root
		path = filepath.Join(gopath, "bin")
		if ok, _ := pathexists(filepath.Join(path, cfgfile)); ok {
			ConfigPath = path
			return
		}
		// try to get from source code
		path = filepath.Join(gopath, "src/github.com/schwarzlichtbezirk/pds/config")
		if ok, _ := pathexists(path); ok {
			ConfigPath = path
			return
		}
	}

	// no config was found
	log.Fatal("no configuration path was found")
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
