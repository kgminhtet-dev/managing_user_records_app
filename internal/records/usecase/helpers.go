package usecase

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"time"
)

func NewRecord(event string, payload *mqueue.Payload) *data.Record {
	return &data.Record{
		UserID:    payload.UserID,
		Event:     event,
		Data:      payload.Data,
		Timestamp: time.Now(),
	}
}
