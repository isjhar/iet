package pkg

import (
	"errors"
	"os"
	"path/filepath"
)

func FindProjectRoot() (string, error) {
	targetFile := "go.mod"
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		// Check if the target file exists in the current directory
		configPath := filepath.Join(dir, targetFile)
		if _, err := os.Stat(configPath); err == nil {
			return dir, nil // Found it
		}

		// Go up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			break // Reached root, stop
		}
		dir = parent
	}

	return "", errors.New("config file not found in any parent directory")
}
