package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-co-op/gocron/v2"
)

type Cronic struct {
	Path      string
	Config    Config
	Scheduler gocron.Scheduler
	Server    *http.Server
}

func NewCronic() Cronic {
	cronic := Cronic{}
	if len(os.Args) >= 2 {
		// Path to job directory was specified
		path, err := filepath.Abs(os.Args[1])
		if err != nil {
			panic(err)
		}
		info, err := os.Stat(path)
		if err != nil {
			panic(err)
		}
		if !info.IsDir() {
			panic(fmt.Errorf("Path is not a directory: %s", path))
		}
		cronic.Path = path
	} else {
		path, err := filepath.Abs(".")
		if err != nil {
			panic(err)
		}
		cronic.Path = path
	}
	if err := os.Chdir(cronic.Path); err != nil {
		panic(err)
	}
	cronic.Config = LoadConfig()
	cronic.Scheduler = LoadScheduler(cronic.Config)
	cronic.Server = NewServer(cronic.Config)
	return cronic
}
