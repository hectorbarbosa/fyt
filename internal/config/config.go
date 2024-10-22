package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type loggerConfig struct {
	LoggingLevel int `yaml:"logging-level"`
}

type manticoreConfig struct {
	ManticoreURL string `yaml:"manticore-url"`
}

type serverConfig struct {
	ServerURL string `yaml:"server-url"`
}

type postgresConfig struct {
	PostgresURL string `yaml:"postgres-url"`
}

type Config struct {
	Logger    loggerConfig    `yaml:"logger-config"`
	Manticore manticoreConfig `yaml:"manticore-config"`
	Server    serverConfig    `yaml:"server-config"`
	Postgres  postgresConfig  `yaml:"postgres-config"`
}

func NewConfig(configPath string) (*Config, error) {
	var cfg Config

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return &Config{}, err
	}

	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return &Config{}, err
	}
	return &cfg, nil
}
