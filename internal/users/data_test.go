package users

import (
	"os"
	"testing"
)

func TestNewDatabase(t *testing.T) {
	os.Setenv("env", "development")
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatal("Expected error to be nil, but got", err)
	}

	db := newDatabase(&cfg.Database)
	if db.Name() != "postgres" {
		t.Errorf("Expected database name to be %s, but got %s", "postgres", db.Name())
	}
}
