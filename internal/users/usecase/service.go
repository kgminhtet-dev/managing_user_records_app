package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service struct {
	repository *repository.Repository
}

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrEmailAlreadyExist = errors.New("email already exits")
	ErrInternal          = errors.New("internal server error")
)

func (s *Service) CreateUser(user *data.User) error {
	if fetchedUser, _ := s.repository.FindByEmail(user.Email); fetchedUser != nil {
		return ErrEmailAlreadyExist
	}

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		return ErrInternal
	}

	user.ID = uuid.New().String()
	user.Password = string(password)
	if err := s.repository.Create(user); err != nil {
		return ErrInternal
	}

	return nil
}

func (s *Service) GetUsers(page, limit int) ([]*data.User, error) {
	start := (page - 1) * limit
	if start <= 0 {
		start = 1
	}
	end := start + limit

	users, err := s.repository.GetAll(start, end)
	if err != nil {
		return nil, ErrInternal
	}
	
	return users, nil
}

func (s *Service) GetUserById(id string) (*data.User, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, ErrInternal
	}

	return user, nil
}

func (s *Service) UpdateUser(id string, user *data.User) error {
	gotUser, _ := s.repository.GetById(id)
	if gotUser == nil {
		return ErrUserNotFound
	}

	err := s.repository.Update(id, user)
	if err != nil {
		return ErrInternal
	}

	return nil
}

func (s *Service) DeleteUser(id string) error {
	gotUser, _ := s.repository.GetById(id)
	if gotUser == nil {
		return ErrUserNotFound
	}

	err := s.repository.Delete(id)
	if err != nil {
		return ErrInternal
	}

	return nil
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}
