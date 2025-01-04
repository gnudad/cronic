package main

import "os"
import "github.com/goccy/go-yaml"

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

func LoadConfig() Config {
	var config Config
	data, err := os.ReadFile("cronic.yaml")
	if err != nil {
		panic(err)
	}
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
	return config
}