package main

import (
	"context"
	"errors"
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

func (s *Server) Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		_, err := io.WriteString(w, "Cronic Scheulder")
		if err != nil {
			panic(err)
		}
	})
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.Host, s.Port),
		Handler: mux,
	}
	go func() {
		if err := s.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
		fmt.Println("Shutting down web server...")
	}()
}

func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.httpServer.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Web server shut down gracefully.")
}
