package unit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/xehapa/jago/config"
	"github.com/xehapa/jago/utils"
)

func TestNewConfig(t *testing.T) {
	// Get the project root directory
	projectRoot := utils.GetProjectRoot()
	if projectRoot == "" {
		t.Fatal("Failed to determine project root directory")
	}

	// Construct the absolute file path to the .env.test file
	envFile := filepath.Join(projectRoot, ".env.test")

	cfg := config.NewConfig(&envFile)

	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if cfg.ClientId != clientId {
		t.Errorf("Expected ClientId to be %s, but got %s", clientId, cfg.ClientId)
	}

	if cfg.ClientSecret != clientSecret {
		t.Errorf("Expected ClientSecret to be %s, but got %s", clientSecret, cfg.ClientSecret)
	}
}
