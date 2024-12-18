package users

import (
	"os"
	"testing"
)

func TestService(t *testing.T) {
	os.Setenv("env", "development")
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("error in loading config %v", err)
	}
	db := newDatabase(&cfg.Database)
	repo := newRepository(db)
	service := newService(repo)

	t.Run("Create user service", func(t *testing.T) {
		testcases := []struct {
			input *User
			err   error
		}{
			{
				input: &User{Name: "username one", Email: "username1@example.com", Password: "12345678"},
				err:   nil,
			},
			{
				input: &User{Name: "username one", Email: "username1@example.com", Password: "12345678"},
				err:   ErrEmailAlreadyExist,
			},
		}

		for _, tc := range testcases {
			resultErr := service.CreateUser(tc.input)
			if resultErr != tc.err {
				t.Errorf("Expected error %v, but got %v", tc.err, resultErr)
			}
		}
	})

}
