package unit

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/xehapa/jago/config"
)

func TestNewConfig(t *testing.T) {
	basePath, err := os.Getwd()

	if err != nil {
		t.Fatal("Error getting base path:", err)
	}
	projectRoot := filepath.Dir(filepath.Dir(basePath))
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
