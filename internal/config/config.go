package config

import (
	"bufio"
	"fmt"
	"io"
	"log"
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

func Load(r io.Reader) (*Config, error) {
	reader := bufio.NewReader(r)
	data := make([]byte, 1<<10)
	n, err := reader.Read(data)
	if err != nil {
		log.Fatalf("Can't read reader: %v", err)
	}

	var configs Configs
	if err := yaml.Unmarshal(data[:n], &configs); err != nil {
		return nil, err
	}

	switch env := os.Getenv("ENV"); env {
	case "development":
		return &configs.Development, nil
	case "testing":
		return &configs.Testing, nil
	default:
		return nil, fmt.Errorf("invalid environment %q", env)
	}
}
