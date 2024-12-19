package config

import (
	"os"
	"strings"
	"testing"
)

const testConfig = `
development:
  database:
    name: postgres
    host: localhost
    port: 5432
    user: postgres
    password: 12345678
    dbname: mur_user
    sslmode: disable
    timezone: Asia/Yangon

testing:
  database:
    name: sqlite
`

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
			if err := os.Setenv("ENV", tc.env); err != nil {
				t.Fatal("Setting environment variable error")
			}

			resultConfig, err := Load(strings.NewReader(testConfig))
			if err != nil {
				t.Fatalf("Load config error: %v", err)
			}

			if *resultConfig.Database != *tc.expectedConfig.Database {
				t.Errorf("Expected config %v, but got %v", *tc.expectedConfig, *resultConfig)
			}
		})
	}
}
