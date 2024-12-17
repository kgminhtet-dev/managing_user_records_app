package config

import (
	"github.com/spf13/viper"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	cfg := Load()

	if cfg.Host != "localhost" {
		t.Error("Expected host to be localhost, got", viper.GetString("HOST"))
	}

	if cfg.Port != "8080" {
		t.Error("Expected port to be 8080, got", viper.GetString("PORT"))
	}
}
