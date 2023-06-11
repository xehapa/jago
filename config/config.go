package config

import (
	"github.com/xehapa/jago/models"
)

type Config struct {
	ClientId     string `json:"apiKey"`
	ClientSecret string `json:"apiSecret"`
}

func NewConfig() *Config {
	return &Config{
		ClientId:     models.ClientId,
		ClientSecret: models.ClientSecret,
	}
}
