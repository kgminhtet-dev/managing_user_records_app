package usecases

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/config"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/repository"
	"log"
	"os"
	"testing"
)

var service *Service
var users []*data.User

func TestMain(m *testing.M) {
	if err := os.Setenv("env", "testing"); err != nil {
		log.Fatal("Error setting environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error loading configuration")
	}

	db := data.New(cfg.Database)
	repo := repository.New(db)
	service = NewService(repo)
	users = generateRandomUsers(10)
	db.CreateInBatches(users, len(users))

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestService_CreateUser(t *testing.T) {
	user := &data.User{
		ID:       uuid.New().String(),
		Name:     "user 000",
		Email:    "usertrplezeor@example.com",
		Password: "12345678",
	}
	testcases := []struct {
		name        string
		user        *data.User
		expectedErr error
	}{
		{
			name:        "Creating a new user",
			user:        user,
			expectedErr: nil,
		},
		{
			name:        "Creating a new user",
			user:        user,
			expectedErr: ErrEmailAlreadyExist,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := service.CreateUser(tc.user)
			if err == nil {
				if tc.expectedErr != nil {
					t.Errorf("Expected error to be nil, but got %v", err)
				}
			} else if !errors.Is(err, ErrEmailAlreadyExist) {
				t.Errorf("Expected error %v, but got %v", ErrEmailAlreadyExist, err)
			}

			fetcheduser, err := service.GetUserById(tc.user.ID)
			if fetcheduser == nil {
				t.Fatalf("Expected created user %v but got %v", tc.user, fetcheduser)
			}

			if fetcheduser.ID != tc.user.ID {
				t.Errorf("Expected user id %s but got %s", tc.user.ID, fetcheduser.ID)
			}
		})
	}
}

func generateRandomUsers(size int) []*data.User {
	users := make([]*data.User, size, size)

	for i, _ := range users {
		user := data.User{}
		user.ID = uuid.New().String()
		user.Name = fmt.Sprintf("user%d", i+1)
		user.Email = fmt.Sprintf("user%d@example.com", i+1)
		users[i] = &user
	}

	return users
}
