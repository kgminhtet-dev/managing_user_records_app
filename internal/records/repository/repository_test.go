package repository_test

import (
	"context"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/repository"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/testutil"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"testing"
	"time"
)

var (
	collection *mongo.Collection
)

func TestMain(m *testing.M) {
	var database *mongo.Database
	database, collection = testutil.NewEnvironment()

	exitCode := m.Run()

	(func() {
		testutil.Clear(context.TODO(), database)
		data.DisconnectDatabase(context.TODO(), database)
	})()

	os.Exit(exitCode)
}

func TestCreate(t *testing.T) {
	testcases := []struct {
		name        string
		record      *data.UserRecord
		expectedErr error
	}{
		{
			name: "Create a new record",
			record: &data.UserRecord{
				ID:        primitive.ObjectID{},
				UserID:    uuid.New().String(),
				Event:     "UserCreated",
				Data:      nil,
				Timestamp: time.Now(),
			},
		},
	}
	repo := repository.New(collection)

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := repo.Create(context.Background(), tc.record)
			assert.Equal(t, err, tc.expectedErr)

			if err == nil {
				var fetchedRecord data.UserRecord
				err = collection.FindOne(context.Background(), bson.M{"user_id": tc.record.UserID}).Decode(&fetchedRecord)

				assert.NoError(t, err)
				assert.Equal(t, fetchedRecord.UserID, tc.record.UserID)
				assert.Equal(t, fetchedRecord.Event, tc.record.Event)
			}
		})
	}
}

func TestGetById(t *testing.T) {
	records := testutil.GenerateRandomRecords(10)
	testutil.SeedDatabase(collection, records)

	repo := repository.New(collection)
	record := records[0].(*data.UserRecord)
	fetchedRecord, err := repo.GetById(context.Background(), record.ID)

	assert.NoError(t, err)
	assert.NotNil(t, fetchedRecord)
	assert.Equal(t, record.ID, fetchedRecord.ID)
}

func TestGetAll(t *testing.T) {
	records := testutil.GenerateRandomRecords(10)
	testutil.SeedDatabase(collection, records)

	repo := repository.New(collection)
	start := 1
	limit := 5
	fetchedRecords, err := repo.GetAll(context.Background(), start, limit)

	assert.NoError(t, err)
	assert.NotNil(t, fetchedRecords)
	assert.Len(t, fetchedRecords, limit)
}
