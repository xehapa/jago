package unit

import (
	"os"
	"testing"

	"github.com/xehapa/jago/config"
)

func TestNewConfig(t *testing.T) {
	os.Setenv("TEST_ENV", "true")
	envFile := "../../.env.test"

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
