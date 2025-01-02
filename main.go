package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/goccy/go-yaml"
)

type Server struct {
    Host string
    Port int
}

type Job struct {
    Name string
    Cron string
    Cmd  string
}

type Config struct {
    Server Server
    Jobs   map[string]Job
}

func main() {
    if len(os.Args) == 2 {
        // Path to job directory was specified
        path := os.Args[1]
        info, err := os.Stat(path)
        if err != nil {
            panic(err)
        }
        if !info.IsDir() {
            panic(fmt.Errorf("Path is not a directory: %s", path))
        }
        os.Chdir(path)
    }

    data, err := os.ReadFile("cronic.yaml")
    if err != nil {
        panic(err)
    }

    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        panic(err)
    }

    yamlData, err := yaml.MarshalWithOptions(&config, yaml.IndentSequence(true))
    if err != nil {
        panic(err)
    }
    if err := os.WriteFile("cronic2.yaml", yamlData, 0600); err != nil {
        panic(err)
    }

    for key, job := range config.Jobs {
        fmt.Printf("%s (%s): %s\n", key, job.Name, job.Cmd)
        cmd := exec.Command("sh", "-c", job.Cmd)
        cmd.Stdout = os.Stdout
        if err := cmd.Run(); err != nil {
            panic(err)
        }
    }
}
