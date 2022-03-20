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
	var c Config
	data, err := os.ReadFile(filepath)
	if err != nil {
		return c, err
	}

	if err != yaml.Unmarshal(data, &c) {
		return c, err
	}

	return c, nil
}
