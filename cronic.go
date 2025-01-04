package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Cronic struct {
	path   string
	config Config
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
		cronic.path = path
	} else {
		path, err := filepath.Abs(".")
		if err != nil {
			panic(err)
		}
		cronic.path = path
	}
	if err := os.Chdir(cronic.path); err != nil {
		panic(err)
	}
	cronic.config = LoadConfig()
	return cronic
}
