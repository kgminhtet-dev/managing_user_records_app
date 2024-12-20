package usecase

import (
	"context"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/repository"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/testutil"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
	"time"
)

var (
	repo       *repository.Repository
	collection *mongo.Collection
)

func TestMain(m *testing.M) {
	var database *mongo.Database
	database, collection = testutil.NewEnvironment()
	repo = repository.New(collection)

	exitCode := m.Run()

	(func() {
		testutil.Clear(database)
		data.DisconnectDatabase(nil, database)
	})()

	os.Exit(exitCode)
}

func TestService_CreateRecord(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	events := []string{"UserCreated", "UserUpdated", "UserDeleted", "UserFetched"}
	testcases := []struct {
		name    string
		payload *Payload
		event   string
	}{
		{
			name: "Create a new record",
			payload: &Payload{
				UserID: uuid.New().String(),
				Data:   "Hello World",
			},
			event: events[0],
		},
		{
			name: "Update a record",
			payload: &Payload{
				UserID: uuid.New().String(),
				Data:   "Hello World",
			},
			event: events[1],
		},
		{
			name: "Delete a record",
			payload: &Payload{
				UserID: uuid.New().String(),
				Data:   "Hello World",
			},
			event: events[2],
		},
		{
			name: "Fetch a record",
			payload: &Payload{
				UserID: uuid.New().String(),
				Data:   "Hello World",
			},
			event: events[3],
		},
	}

	service := NewService(repo)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			msg := &Message{
				Event:   tc.event,
				Payload: tc.payload,
			}

			err := service.CreateRecord(ctx, msg)
			assert.NoError(t, err)

			var record data.UserRecord
			err = collection.FindOne(ctx, bson.M{"user_id": tc.payload.UserID}).Decode(&record)

			assert.NoError(t, err)
			assert.Equal(t, record.Event, msg.Event)
			assert.Equal(t, tc.payload.UserID, record.UserID)
		})
	}
}
