package unit

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/xehapa/jago/config"
)

func TestNewConfig(t *testing.T) {
	// Get the current directory (where the test file is located)
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(filename)

	// Construct the absolute file path to the .env.test file
	envFile := filepath.Join(testDir, "..", "..", ".env.test")

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
