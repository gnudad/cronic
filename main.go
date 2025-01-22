package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	cronic := NewCronic()
	fmt.Println("Running cronic from", cronic.path)
	cronic.Start()

	ctx, quit := signal.NotifyContext(context.Background(), os.Interrupt)
	defer quit()
	<-ctx.Done()

	if err := cronic.Shutdown(); err != nil {
		panic(err)
	}
}
