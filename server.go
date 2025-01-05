package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func NewServer(config Config) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5*time.Second)
		_, err := io.WriteString(w, "Cronic Scheduler")
		if err != nil {
			panic(err)
		}
	})
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: mux,
	}

}
