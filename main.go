package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cronic := NewCronic()
	fmt.Println("Running cronic from", cronic.Path)
	cronic.Scheduler.Start()

	go func() {
		if err := cronic.Server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
		fmt.Println("Shutting down web server...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	err := cronic.Scheduler.Shutdown()
	if err != nil {
		panic(err)
	}

	if err := cronic.Server.Shutdown(shutdownCtx); err != nil {
		panic(fmt.Errorf("HTTP shutdown error: %v", err))
	}
	fmt.Println("Graceful shutdown complete.")
}
