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
	fmt.Println("Running cronic from", cronic.path)
	cronic.scheduler.Start()

	go func() {
		if err := cronic.Server.httpServer.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
		fmt.Println("Shutting down web server...")
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	err := cronic.scheduler.Shutdown()
	if err != nil {
		panic(err)
	}

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()
	if err := cronic.Server.httpServer.Shutdown(shutdownCtx); err != nil {
		panic(fmt.Errorf("HTTP shutdown error: %v", err))
	}

	fmt.Println("Graceful shutdown complete.")
}
