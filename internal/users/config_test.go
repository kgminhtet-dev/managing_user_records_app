package users

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("Test Development Environment", func(t *testing.T) {
		os.Setenv("env", "development")
		expected := &Config{
			Database: DatabaseConfig{
				Host:     "localhost",
				Port:     "5432",
				User:     "postgres",
				Password: "12345678",
				DbName:   "mur_user",
				SSLMode:  "disable",
				TimeZone: "Asia/Yangon",
			},
		}

		result, err := LoadConfig()
		if err != nil {
			t.Error("Expected error to be nil, but got", err)
		}

		if *result != *expected {
			t.Errorf("Expected %v, but got %v", expected, result)
		}
	})
}
