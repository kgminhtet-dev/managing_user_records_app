package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	testcases := []struct {
		name       string
		env        string
		dbName     string
		collection string
	}{
		{
			name:       "Development config",
			env:        "development",
			dbName:     "mur",
			collection: "records",
		},
		{
			name:       "Testing config",
			env:        "testing",
			dbName:     "test",
			collection: "records",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if err := os.Setenv("ENV", tc.env); err != nil {
				t.Fatal("Error setting environment variables:", err)
			}

			cfg := LoadConfig(".")
			assert.Equal(t, tc.dbName, cfg.Database.Name)
			assert.Equal(t, tc.collection, cfg.Database.Collection)
		})
	}
}
