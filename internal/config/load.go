package config

import (
	"encoding/json"
	"os"
)

// Representation of the JSON config file
type Config struct {
	HTTP struct {
		Bind string `json:"bind"`
		Port int    `json:"port"`
	} `json:"http"`
	Services []Service `json:"services"`
}

type Service struct {
	Name string `json:"name"`
	Path string `json:"path"`
	URL  string `json:"url"`
}

// Global instance of Config
var C Config

// Load configuration
func Load() error {
	// Open file
	file, err := os.Open("config.json")
	if err != nil {
		return err
	}

	// Parse JSON file
	err = json.NewDecoder(file).Decode(&C)
	if err != nil {
		return err
	}

	return nil
}
