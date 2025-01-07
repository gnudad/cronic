package main

import (
	"os"

	"github.com/goccy/go-yaml"
)

func LoadConfig(cronic *Cronic) error {
	data, err := os.ReadFile("cronic.yaml")
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(data, cronic); err != nil {
		return err
	}
	yamlData, err := yaml.MarshalWithOptions(cronic, yaml.IndentSequence(true))
	if err != nil {
		return err
	}

	// Write it back to another file for now to see if/how it changes
	if err := os.WriteFile("cronic2.yaml", yamlData, 0600); err != nil {
		return err
	}

	return nil
}
