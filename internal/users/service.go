package users

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	repository *Repository
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExist = errors.New("email already exits")
	ErrInternal          = errors.New("internal server error")
)

func (s *Service) CreateUser(user *User) error {
	if fetchedUser, _ := s.repository.FindByEmail(user.Email); fetchedUser != nil {
		return ErrEmailAlreadyExist
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		return ErrInternal
	}

	user.Password = string(password)
	if err := s.repository.Create(user); err != nil {
		return ErrInternal
	}

	return nil
}

func (s *Service) GetUsers(page, limit int) ([]*User, error) {
	return nil, nil
}

func (s *Service) GetUserById(id string) (*User, error) {
	return nil, nil
}

func (s *Service) UpdateUser(id string, user *User) (*User, error) {
	return nil, nil
}

func (s *Service) DeleteUser(id string) (*User, error) {
	return nil, nil
}

func newService(repo *Repository) *Service {
	return &Service{repository: repo}
}
