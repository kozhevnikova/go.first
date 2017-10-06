package main

import (
	"os"

	"github.com/naoina/toml"
)

func parseConfig() (Config, error) {
	var config Config
	f, err := os.Open("config.toml")
	if err != nil {
		return config, err
	}
	defer f.Close()

	if err := toml.NewDecoder(f).Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}
