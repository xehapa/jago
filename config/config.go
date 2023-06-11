package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ClientId     string
	ClientSecret string
	AuthUrl      string
}

func NewConfig(envFile *string) *Config {
	if envFile == nil {
		env := ".env"
		envFile = &env
	}

	err := godotenv.Load(*envFile)

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	authUrl := os.Getenv("AUTH_URL")

	return &Config{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		AuthUrl:      authUrl,
	}
}
