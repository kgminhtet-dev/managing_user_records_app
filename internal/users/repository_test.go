package users

import (
	"os"
	"testing"
)

func TestRepository(t *testing.T) {
	os.Setenv("env", "development")
	cfg, _ := LoadConfig()
	db := newDatabase(&cfg.Database)
	repo := newRepository(db)

	t.Run("Create user", func(t *testing.T) {
		user := &User{
			Name:     "user1",
			Email:    "user1@example.com",
			Password: "123456768",
		}

		createdUser, err := repo.Create(user)
		if err != nil {
			t.Errorf("Expected error to be nil, got %v", err)
		}

		if createdUser == nil {
			t.Fatalf("Expected user not to be nil")
		}

		if createdUser.Email != user.Email {
			t.Errorf("Expected new user email %v, but got %v", user.Email, createdUser.Email)
		}
	})
}
