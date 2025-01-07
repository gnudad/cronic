package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type Server struct {
	Host       string
	Port       int
	httpServer *http.Server
}

func InitServer(cronic *Cronic) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		_, err := io.WriteString(w, "Cronic Scheduler")
		if err != nil {
			panic(err)
		}
	})
	cronic.Server.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cronic.Server.Host, cronic.Server.Port),
		Handler: mux,
	}
	return nil
}
