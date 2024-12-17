package users

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Configs struct {
	Development Config `yaml:"development"`
	Production  Config `yaml:"production"`
}

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DbName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
	TimeZone string `yaml:"timezone"`
}

func LoadConfig() (*Config, error) {
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
	default:
		return nil, fmt.Errorf("invalid environment %q", env)
	}
}
