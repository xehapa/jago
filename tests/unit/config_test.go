package unit

import (
	"testing"

	"github.com/xehapa/jago/config"
	"github.com/xehapa/jago/models"
)

func TestNewConfig(t *testing.T) {
	clientId := models.ClientId
	clientSecret := models.ClientSecret

	cfg := config.NewConfig()

	if cfg.ClientId != clientId {
		t.Errorf("Expected ClientId to be %s, but got %s", clientId, cfg.ClientId)
	}

	if cfg.ClientSecret != clientSecret {
		t.Errorf("Expected ClientSecret to be %s, but got %s", clientSecret, cfg.ClientSecret)
	}
}
