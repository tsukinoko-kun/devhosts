package main

import (
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

var defaultConfigFiles = []string{"devhosts.yaml", "devhosts.yml", "devhosts.toml", "devhosts.json", "devhosts.jsonc", "devhosts.json5"}

func LoadConfig(path string) (map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cfg := make(map[string]string)

	switch filepath.Ext(path) {
	case ".json", ".jsonc", ".json5":
		err = json.NewDecoder(f).Decode(&cfg)
	case ".yaml", ".yml":
		err = yaml.NewDecoder(f).Decode(&cfg)
	case ".toml":
		_, err = toml.NewDecoder(f).Decode(&cfg)
	default:
		err = errors.New("unsupported config file format")
	}

	return cfg, err
}

func FindConfig() (string, error) {
	for _, name := range defaultConfigFiles {
		if _, err := os.Stat(name); err == nil {
			return name, nil
		}
	}

	return "", errors.New("failed to find config file")
}
