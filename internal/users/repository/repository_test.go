package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/testutil"
	"github.com/mattn/go-sqlite3"
	"os"
	"testing"
)

var (
	repo  *Repository
	users []*data.User
)

func TestMain(m *testing.M) {
	db := testutil.Setup()
	users = testutil.GenerateRandomUsers(10)
	testutil.SeedDatabase(db, users)
	repo = New(db)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreateUser(t *testing.T) {
	user := &data.User{
		ID:       uuid.New().String(),
		Name:     "user14",
		Email:    "user14@example.com",
		Password: "123456768",
	}

	testcases := []struct {
		name           string
		user           *data.User
		expectedError  error
		expectedResult *data.User
	}{
		{
			name:          "Create a new user",
			user:          user,
			expectedError: nil,
		},
		{
			name:          "Create existing user",
			user:          user,
			expectedError: sqlite3.ErrConstraintUnique,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Create(tc.user)
			if err == nil {
				if err != tc.expectedError {
					t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
				}
			} else if errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error %v, but got %v", tc.expectedError, err)
			}
		})
	}
}

func TestGetUserById(t *testing.T) {
	user := users[0]
	fetchedUser, err := repo.GetById(user.ID)
	if err != nil {
		t.Fatalf("Expected error to be nil.")
	}

	if fetchedUser.ID != user.ID {
		t.Errorf("Expected user id %s, but got %s", user.ID, fetchedUser.ID)
	}
}

func TestFindUserByEmail(t *testing.T) {
	user := users[0]
	fetchedUser, err := repo.FindByEmail(user.Email)
	if err != nil {
		t.Fatalf("Expected error to be nil.")
	}

	if fetchedUser.Email != user.Email {
		t.Errorf("Expected user id %s, but got %s", user.Email, fetchedUser.Email)
	}
}

func TestGetUsers(t *testing.T) {
	start, end := 1, 10
	fetchedusers, err := repo.GetAll(start, end)
	if err != nil {
		t.Fatalf("Error found in getting user %v", err)
	}

	if len(fetchedusers) != (end - start) {
		t.Errorf("Expected length %d, but got %d", (end - start), len(fetchedusers))
	}
}

func TestUpdateUser(t *testing.T) {
	user := users[0]
	updateUserData := &data.User{Email: fmt.Sprintf("updated%s", user.Email)}
	if err := repo.Update(user.ID, updateUserData); err != nil {
		t.Fatalf("Error found in updating %v", err)
	}

	user, err := repo.GetById(user.ID)
	if err != nil {
		t.Fatalf("Error found in get by id %v", err)
	}

	if user.Email != updateUserData.Email {
		t.Errorf("Expected email %s, but got %s", updateUserData.Email, user.Email)
	}
}

func TestDeleteUser(t *testing.T) {
	user := users[0]
	if err := repo.Delete(user.ID); err != nil {
		t.Fatalf("Error soft deleting user %v", err)
	}

	if fetchedUser, _ := repo.GetById(user.ID); fetchedUser != nil {
		t.Errorf("Expected user to be nil, but got %v", fetchedUser)
	}
}
