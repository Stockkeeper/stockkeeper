package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	ServerHost string `yaml:"host"`
	ServerPort string `yaml:"port"`
}

func ParseConfig(filepath string) (Config, error) {
	var cfg Config
	data, err := os.ReadFile(filepath)
	if err != nil {
		return cfg, err
	}

	if err != yaml.Unmarshal(data, &cfg) {
		return cfg, err
	}

	return cfg, nil
}
