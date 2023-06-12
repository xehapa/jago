package utils

import (
	"log"
	"os"
	"path/filepath"
)

func GetProjectRoot() string {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir
		}

		parent := filepath.Dir(dir)

		if parent == dir {
			break
		}

		dir = parent
	}

	return ""
}
