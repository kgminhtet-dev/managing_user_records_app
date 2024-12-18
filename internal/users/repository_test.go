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
	user := &User{
		Name:     "user14",
		Email:    "user14@example.com",
		Password: "123456768",
	}
	updateUser := &User{
		Email: "user144@example.com",
	}

	t.Run("Create user", func(t *testing.T) {
		err := repo.Create(user)
		if err != nil {
			t.Fatalf("Expected error to be nil, got %v", err)
		}

		var fetchedUser User
		db.First(&fetchedUser, "email = ?", user.Email)

		if fetchedUser.ID != user.ID {
			t.Errorf("Expected user id %s, but got id %s", user.ID, fetchedUser.ID)
		}

		*user = fetchedUser
	})

	t.Run("Get all users", func(t *testing.T) {
		start, end := 1, 5
		limit := end - start

		users, err := repo.GetAll(start, end)
		if err != nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}

		if len(users) != limit {
			t.Errorf("Expected user count %d, but got %d", limit, len(users))
		}
	})

	t.Run("Get user by id", func(t *testing.T) {
		id := user.ID.String()
		user, err := repo.GetById(id)
		if err != nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}

		if user.ID.String() != id {
			t.Errorf("Expected id %s, but got %s", id, user.ID)
		}
	})

	t.Run("Update user's email", func(t *testing.T) {
		if err := repo.Update(user.ID.String(), updateUser); err != nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}

		fetchedUser, _ := repo.GetById(user.ID.String())
		if fetchedUser.Email != updateUser.Email {
			t.Errorf("Expected email %s, but got %s", user.Email, fetchedUser.Email)
		}

		user = fetchedUser
	})

	t.Run("Find user by email", func(t *testing.T) {
		fetchedUser, err := repo.FindByEmail(user.Email)
		if err != nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}

		if fetchedUser.Email != user.Email {
			t.Errorf("Expected email %s, but got %s", user.Email, fetchedUser.Email)
		}
	})

	t.Run("Soft delete user", func(t *testing.T) {
		if err := repo.Delete(user.ID.String()); err != nil {
			t.Fatalf("Expected error to be nil, but got %v", err)
		}

		if _, err := repo.GetById(user.ID.String()); err == nil {
			t.Errorf("Expected error to be nil, but got %v", err)
		}
	})
}
