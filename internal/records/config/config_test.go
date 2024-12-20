package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	testcases := []struct {
		name   string
		env    string
		dbName string
	}{
		{
			name:   "Development config",
			env:    "development",
			dbName: "mur_records",
		},
		{
			name:   "Testing config",
			env:    "testing",
			dbName: "test",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if err := os.Setenv("ENV", tc.env); err != nil {
				t.Fatal("Error setting environment variables:", err)
			}

			cfg := LoadConfig(".")
			if cfg.Database.Name != tc.dbName {
				t.Errorf("Expected database name to be %s, but got %s", tc.dbName, cfg.Database.Name)
			}
		})
	}
}
