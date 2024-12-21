package usecase

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/repository"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/testutil"
	"os"
	"testing"
)

var (
	service *Service
	users   []*data.User
)

func TestMain(m *testing.M) {
	db := testutil.Setup()
	users = testutil.GenerateRandomUsers(10)
	testutil.SeedDatabase(db, users)
	repo := repository.New(db)
	service = NewService(repo)

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

			fetcheduser, _ := service.GetUserById(tc.user.ID)
			if fetcheduser == nil {
				t.Fatalf("Expected created user %v but got %v", tc.user, fetcheduser)
			}

			if fetcheduser.ID != tc.user.ID {
				t.Errorf("Expected user id %s but got %s", tc.user.ID, fetcheduser.ID)
			}
		})
	}
}

func TestService_GetUserById(t *testing.T) {
	user := users[0]
	testcases := []struct {
		name         string
		id           string
		expectedErr  error
		expectedUser *data.User
	}{
		{
			name:         "Find user by existing id",
			id:           user.ID,
			expectedErr:  nil,
			expectedUser: user,
		},
		{
			name:         "Find user by new id",
			id:           uuid.New().String(),
			expectedErr:  ErrUserNotFound,
			expectedUser: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			gotUser, err := service.GetUserById(tc.id)

			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
			}

			if gotUser != nil && gotUser.ID != tc.expectedUser.ID {
				t.Errorf("Expected user %+v, but got %+v", tc.expectedUser, gotUser)
			}
		})
	}
}

func TestService_GetUsers(t *testing.T) {
	testcases := []struct {
		name  string
		page  int
		limit int
	}{
		{
			name:  "Get all users with positive page number",
			page:  1,
			limit: 10,
		},
		{
			name:  "Get all users with negative page number",
			page:  -1,
			limit: 10,
		},
		{
			name:  "Get all users with zero page number",
			page:  0,
			limit: 10,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := service.GetUsers(tc.page, tc.limit)
			if err != nil {
				t.Errorf("Expected error to be nil, but got %v", err)
			}
		})
	}
}

func TestService_UpdateUser(t *testing.T) {
	oldUser := users[0]
	updateUserInfo := &data.User{Email: fmt.Sprintf("updated%s", oldUser.Email)}

	testcases := []struct {
		name           string
		id             string
		updateUserInfo *data.User
		expectedErr    error
	}{
		{
			name:           "Updating the old user",
			id:             oldUser.ID,
			updateUserInfo: updateUserInfo,
			expectedErr:    nil,
		},
		{
			name:           "Updating the not existed user",
			id:             uuid.New().String(),
			updateUserInfo: updateUserInfo,
			expectedErr:    ErrUserNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := service.UpdateUser(tc.id, tc.updateUserInfo)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expcted error %v, but got %v", tc.expectedErr, err)
			}
		})
	}
}

func TestService_DeleteUser(t *testing.T) {
	user := users[0]
	testcases := []struct {
		name        string
		id          string
		expectedErr error
	}{
		{
			name:        "Deleting the existing user",
			id:          user.ID,
			expectedErr: nil,
		},
		{
			name:        "Deleting the deleted user",
			id:          user.ID,
			expectedErr: ErrUserNotFound,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := service.DeleteUser(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("Expected error %v but got %v", tc.expectedErr, err)
			}
		})
	}
}
