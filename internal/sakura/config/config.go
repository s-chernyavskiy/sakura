package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Port       string `yaml:"host"`
	Host       int    `yaml:"port"`
	MaxClients int    `yaml:"max-clients"`
	MaxTimeout int `yaml:"max-timeout"`
}

func NewConfig() *AppConfig {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		panic(err)
	}

	var config AppConfig

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
