package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Database struct {
	Url  string
	Name string
}

type Config struct {
	Database Database
}

func LoadConfig(path string) *Config {
	env := os.Getenv("ENV")

	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %v", err)
	}

	var config Config
	url := viper.GetString(fmt.Sprintf("%s.database.url", env))
	name := viper.GetString(fmt.Sprintf("%s.database.name", env))

	config.Database.Url = url
	config.Database.Name = name

	return &config
}
