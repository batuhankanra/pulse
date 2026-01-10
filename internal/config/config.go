package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/batuhankanra/pulse.git/internal/models"
)

func LoadConfig() (*models.Config, error) {
	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".pulse.json")

	cfg := &models.Config{
		URLs:    map[string]string{},
		Headers: map[string]string{},
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return cfg, nil
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, cfg)
	return cfg, nil
}
