package usecase

import (
	"context"
	"errors"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/repository"
)

var (
	ErrNotFound       = errors.New("record not found")
	ErrInvalidPayload = errors.New("invalid payload")
	ErrInternalServer = errors.New("internal server error")
	ErrDatabaseError  = errors.New("database error")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrBadRequest     = errors.New("bad request")
)

type Message struct {
	Event   string
	Payload any
}

type Service struct {
	repository *repository.Repository
}

func (s *Service) CreateRecord(ctx context.Context, msg *Message) error {
	payload, ok := msg.Payload.(*Payload)

	if !ok || !payload.Validate() {
		return ErrInvalidPayload
	}

	record := NewRecord(msg.Event, payload)
	err := s.repository.Create(ctx, record)
	if err != nil {
		return ErrDatabaseError
	}

	return nil
}

func (s *Service) GetRecords(ctx context.Context, page int) ([]*data.UserRecord, error) {
	if page <= 0 {
		page = 1
	}

	limit := 10
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
