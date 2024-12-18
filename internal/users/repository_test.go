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
	user := User{
		Name:     "user5",
		Email:    "user5@example.com",
		Password: "123456768",
	}

	t.Run("Create user", func(t *testing.T) {
		err := repo.Create(&user)
		if err != nil {
			t.Errorf("Expected error to be nil, got %v", err)
		}

		var fetchUser User
		db.First(&fetchUser, "email = ?", user.Email)

		if fetchUser.ID != user.ID {
			t.Errorf("Expected user id %v, but got id %v", user, fetchUser)
		}
	})

	t.Run("Get all users", func(t *testing.T) {
		start, end := 1, 5
		users, err := repo.GetAll(start, end)
		limit := end - start

		if err != nil {
			t.Errorf("Expected error to be nil, but got %v", err)
		}

		if len(users) != limit {
			t.Errorf("Expected user count %d, but got %d", limit, len(users))
		}
	})

	t.Run("Get user by id", func(t *testing.T) {
		id := "5ea5b063-4076-4e02-8650-7d1da3833bf7"
		user, err := repo.GetById(id)
		if err != nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}

		if user.ID.String() != id {
			t.Errorf("Expected id %s, but got %s", id, user.ID)
		}
	})

	t.Run("Find by email", func(t *testing.T) {
		fetchedUser, err := repo.FindByEmail(user.Email)
		if err != nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}

		if fetchedUser.Email != user.Email {
			t.Errorf("Expected email %s, but got %s", user.Email, fetchedUser.Email)
		}
	})
}
