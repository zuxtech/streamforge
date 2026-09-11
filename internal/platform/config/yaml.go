package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

func loadYAML(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}

		return err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return err
	}

	return nil
}
