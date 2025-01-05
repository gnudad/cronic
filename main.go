package main

import (
	"fmt"
	"log"
)

func main() {
	cronic := NewCronic()
	fmt.Println("Running cronic from", cronic.Path)
	cronic.Scheduler.Start()

	log.Fatal(cronic.Server.ListenAndServe())

	err := cronic.Scheduler.Shutdown()
	if err != nil {
		panic(err)
	}
}
