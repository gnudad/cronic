package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cronic := NewCronic()
	fmt.Println("Running cronic from ", cronic.path)

	for key, job := range cronic.config.Jobs {
		fmt.Printf("%s (%s): %s\n", key, job.Name, job.Cmd)
		cmd := exec.Command("sh", "-c", job.Cmd)
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}
}
