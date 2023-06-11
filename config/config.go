package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/xehapa/jago/utils"
)

type Config struct {
	ClientId     string
	ClientSecret string
	AuthUrl      string
}

func NewConfig(isTest bool) *Config {
	projectRoot := utils.GetProjectRoot()
	if projectRoot == "" {
		log.Fatal("Failed to determine project root directory")
	}

	envFile := ".env"

	if isTest {
		envFile = envFile + ".test"
	}

	env := filepath.Join(projectRoot, envFile)
	err := godotenv.Load(env)

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
