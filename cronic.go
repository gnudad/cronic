package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/labstack/echo/v4"
)

type Cronic struct {
	Config    Config
	Jobs      map[string]Job
	path      string
	scheduler gocron.Scheduler
	server    *echo.Echo
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
	if err := LoadConfig(&cronic); err != nil {
		panic(err)
	}
	cronic.scheduler = NewScheduler(&cronic)
	cronic.server = NewServer(&cronic)
	return cronic
}

func (cronic *Cronic) Start() {
	cronic.scheduler.Start()
	go func() {
		err := cronic.server.Start(fmt.Sprintf(":%d", cronic.Config.Port))
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		fmt.Print("\r") // Overwrite ^C
		fmt.Println("Shutting down web server...")
	}()

}

func (cronic *Cronic) Shutdown() error {
	schedulerError := cronic.scheduler.Shutdown()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	serverError := cronic.server.Shutdown(ctx)
	if serverError == nil {
		fmt.Println("Web server shut down gracefully.")
	}

	return errors.Join(schedulerError, serverError)
}
