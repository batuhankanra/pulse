package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/batuhankanra/pulse.git/internal/models"
)

func path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".pulse-config.json"), nil
}

func Load() (*models.Config, error) {
	p, err := path()
	if err != nil {
		return nil, err
	}
	cfg := &models.Config{
		URLs:    map[string]string{},
		Headers: map[string]map[string]string{},
		Body:    map[string]map[string]string{},
	}
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return cfg, err
	}
	b, err := os.ReadFile(p)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(b, cfg)
	return cfg, nil
}

func Save(cfg *models.Config) error {
	p, err := path()
	if err != nil {
		return err
	}
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(p, b, 0644)
}
