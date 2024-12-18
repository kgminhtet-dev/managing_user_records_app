package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"timezone"`
}

type Config struct {
	Database *DatabaseConfig
}

type Configs struct {
	Development Config `yaml:"development"`
	Production  Config `yaml:"production"`
	Testing     Config `yaml:"testing"`
}

func Load() (*Config, error) {
	data, err := os.ReadFile("./config.yaml")
	if err != nil {
		return nil, err
	}

	var configs Configs
	if err := yaml.Unmarshal(data, &configs); err != nil {
		return nil, err
	}

	switch env := os.Getenv("env"); env {
	case "development":
		return &configs.Development, nil
	case "testing":
		return &configs.Testing, nil
	default:
		return nil, fmt.Errorf("invalid environment %q", env)
	}
}
