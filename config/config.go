package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	APIKey    string `json:"apiKey"`
	APISecret string `json:"apiSecret"`
}

func LoadConfig(filename string) *Config {
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(&config)

	if err != nil {
		log.Fatal(err)
	}

	return config
}
