package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Host string
	Port string
	Env  string
}

func Load() *Config {
	viper.SetConfigFile("development.env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("error reading config file: %w", err))
	}

	return &Config{Host: viper.GetString("HOST"), Port: viper.GetString("PORT"), Env: viper.GetString("ENV")}

}
