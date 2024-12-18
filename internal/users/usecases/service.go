package usecases

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
	return nil, nil
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

func (s *Service) UpdateUser(id string, user *data.User) (*data.User, error) {
	return nil, nil
}

func (s *Service) DeleteUser(id string) (*data.User, error) {
	return nil, nil
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}
