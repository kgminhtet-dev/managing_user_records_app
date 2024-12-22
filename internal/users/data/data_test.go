package data

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/config"
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

func TestNewDatabase(t *testing.T) {
	testcases := []struct {
		Env    string
		Dbname string
	}{
		{Env: "development", Dbname: "postgres"},
		{Env: "testing", Dbname: "sqlite"},
	}

	for _, tc := range testcases {
		t.Run(tc.Env, func(t *testing.T) {
			if err := os.Setenv("ENV", tc.Env); err != nil {
				t.Fatal("Setting environment variable error")
			}

			cfg, err := config.Load(strings.NewReader(testConfig))
			if err != nil {
				t.Fatal("Error loading configuration", err)
			}

			db := New(cfg.Database)
			if db.Name() != tc.Dbname {
				t.Errorf("Expected database name to be %s, but got %s", tc.Dbname, db.Name())
			}
		})
	}
}
