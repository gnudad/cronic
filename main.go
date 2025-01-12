package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cronic := NewCronic()
	fmt.Println("Running cronic from", cronic.path)
	cronic.scheduler.Start()
	cronic.Server.Start()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	if err := cronic.scheduler.Shutdown(); err != nil {
		panic(err)
	}
	cronic.Server.Shutdown()
}
