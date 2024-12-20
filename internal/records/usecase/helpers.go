package usecase

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"time"
)

func NewRecord(event string, payload *Payload) *data.UserRecord {
	return &data.UserRecord{
		UserID:    payload.UserID,
		Event:     event,
		Data:      payload.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
