package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	testcases := []struct {
		name           string
		env            string
		expectedConfig *Config
	}{
		{
			name: "Development env config",
			env:  "development",
			expectedConfig: &Config{Database: &DatabaseConfig{
				Name:     "postgres",
				Host:     "localhost",
				Port:     "5432",
				User:     "postgres",
				Password: "12345678",
				DbName:   "mur_user",
				SSLMode:  "disable",
				TimeZone: "Asia/Yangon",
			}}},
		{
			name: "Testing env config",
			env:  "testing",
			expectedConfig: &Config{Database: &DatabaseConfig{
				Name: "sqlite",
			}}},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if err := os.Setenv("env", tc.env); err != nil {
				t.Fatal("Setting environment variable error")
			}

			resultConfig, err := Load()
			if err != nil {
				t.Fatalf("Load config error: %v", err)
			}

			if *resultConfig.Database != *tc.expectedConfig.Database {
				t.Errorf("Expected config %v, but got %v", *tc.expectedConfig, *resultConfig)
			}
		})
	}
}
