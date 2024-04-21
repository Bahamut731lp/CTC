package main

import (
	"os"

	"gopkg.in/yaml.v2"
)

func getConfigFromFile(path string) *Config {
	config := Config{}
	yamlFile, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &config)

	if err != nil {
		panic(err)
	}

	return &config
}
