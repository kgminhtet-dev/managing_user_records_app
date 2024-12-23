package usecase

import (
	"context"
	"errors"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/repository"
	"log"
)

var (
	ErrNotFound       = errors.New("record not found")
	ErrInvalidPayload = errors.New("invalid payload")
	ErrInternalServer = errors.New("internal server error")
	ErrDatabaseError  = errors.New("database error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrBadRequest     = errors.New("bad request")
)

type Service struct {
	repository *repository.Repository
}

func (s *Service) CreateRecord(ctx context.Context, event string, payload *mqueue.Payload) error {
	record := NewRecord(event, payload)
	err := s.repository.Create(ctx, record)
	if err != nil {
		log.Println(err)
		return ErrDatabaseError
	}

	return nil
}

func (s *Service) GetRecords(ctx context.Context, page int, limit int) ([]*data.Record, error) {
	if page <= 0 {
		page = 1
	}

	start := (page - 1) * limit
	records, err := s.repository.GetAll(ctx, start, limit)
	if err != nil {
		return nil, ErrDatabaseError
	}

	if len(records) == 0 {
		return nil, ErrNotFound
	}

	return records, nil
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}
