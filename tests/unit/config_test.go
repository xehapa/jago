package unit

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/xehap/jago/config"
)

func TestNewConfig(t *testing.T) {
	apiKey := "testAPIKey"
	apiSecret := "testAPISecret"

	cfg := config.NewConfig(apiKey, apiSecret)

	if cfg.APIKey != apiKey {
		t.Errorf("Expected APIKey to be %s, but got %s", apiKey, cfg.APIKey)
	}

	if cfg.APISecret != apiSecret {
		t.Errorf("Expected APISecret to be %s, but got %s", apiSecret, cfg.APISecret)
	}
}

func TestLoadConfig(t *testing.T) {
	expectedAPIKey := "testAPIKey"
	expectedAPISecret := "testAPISecret"

	// Create a temporary config file with test data
	tempFile, err := os.CreateTemp("", "config-test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tempFile.Name())

	// Write the test data to the temporary file
	data := []byte(`{
		"apiKey": "` + expectedAPIKey + `",
		"apiSecret": "` + expectedAPISecret + `"
	}`)
	err = ioutil.WriteFile(tempFile.Name(), data, 0644)
	if err != nil {
		t.Fatal(err)
	}

	// Load the config from the temporary file
	cfg := config.LoadConfig(tempFile.Name())

	if cfg.APIKey != expectedAPIKey {
		t.Errorf("Expected APIKey to be %s, but got %s", expectedAPIKey, cfg.APIKey)
	}

	if cfg.APISecret != expectedAPISecret {
		t.Errorf("Expected APISecret to be %s, but got %s", expectedAPISecret, cfg.APISecret)
	}
}
