package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/repository"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreate(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		record := &data.UserRecord{
			UserID:    uuid.New().String(),
			Event:     "create user",
			Data:      nil,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.Create(ctx, record)
		assert.NoError(t, err)
	})

	mt.Run("DatabaseError", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "mock database error",
		}))

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		record := &data.UserRecord{
			UserID:    uuid.New().String(),
			Event:     "create user",
			Data:      nil,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		err := repo.Create(ctx, record)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to insert record")
	})
}

func TestGetById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success", func(mt *mtest.T) {
		expectedID := primitive.NewObjectID()
		mockRecord := bson.D{
			{Key: "_id", Value: expectedID},
			{Key: "user_id", Value: uuid.New().String()},
			{Key: "event", Value: "create user"},
			{Key: "data", Value: nil},
			{Key: "created_at", Value: time.Now()},
			{Key: "updated_at", Value: time.Now()},
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.collection", mtest.FirstBatch, mockRecord))

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		result, err := repo.GetById(ctx, expectedID.Hex())

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, expectedID, result.ID)
		assert.Equal(t, "create user", result.Event)
	})

	mt.Run("NotFound", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: mongo.ErrNoDocuments.Error(),
		}))

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		nonExistentID := primitive.NewObjectID()
		result, err := repo.GetById(ctx, nonExistentID.Hex())

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "no documents")
	})

	mt.Run("DatabaseError", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "mock database error",
		}))

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		randomID := primitive.NewObjectID()
		result, err := repo.GetById(ctx, randomID.Hex())

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Contains(t, err.Error(), "failed to get record")
	})
}

func TestGetAll(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success", func(mt *mtest.T) {
		mockRecords := []bson.D{
			{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "user_id", Value: uuid.New().String()},
				{Key: "event", Value: "create user"},
				{Key: "data", Value: nil},
				{Key: "created_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
			},
			{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "user_id", Value: uuid.New().String()},
				{Key: "event", Value: "update user"},
				{Key: "data", Value: nil},
				{Key: "created_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
			},
		}

		firstBatch := mtest.CreateCursorResponse(1, "test.collection", mtest.FirstBatch, mockRecords...)
		killCursors := mtest.CreateCursorResponse(0, "test.collection", mtest.NextBatch)
		mt.AddMockResponses(firstBatch, killCursors)

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := repo.GetAll(ctx, 0, 2)

		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, "create user", results[0].Event)
		assert.Equal(t, "update user", results[1].Event)
	})

	mt.Run("DatabaseError", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
			Code:    11000,
			Message: "mock database error",
		}))

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := repo.GetAll(ctx, 0, 2)

		assert.Error(t, err)
		assert.Nil(t, results)
		assert.Contains(t, err.Error(), "mock database error")
	})

	mt.Run("EmptyResult", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "test.collection", mtest.FirstBatch))

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := repo.GetAll(ctx, 0, 2)

		assert.NoError(t, err)
		assert.Empty(t, results)
	})

	mt.Run("PaginationTest", func(mt *mtest.T) {
		mockRecords := []bson.D{
			{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "user_id", Value: uuid.New().String()},
				{Key: "event", Value: "create user"},
				{Key: "data", Value: nil},
				{Key: "created_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
			},
			{
				{Key: "_id", Value: primitive.NewObjectID()},
				{Key: "user_id", Value: uuid.New().String()},
				{Key: "event", Value: "update user"},
				{Key: "data", Value: nil},
				{Key: "created_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
			},
		}

		firstBatch := mtest.CreateCursorResponse(1, "test.collection", mtest.FirstBatch, mockRecords...)
		killCursors := mtest.CreateCursorResponse(0, "test.collection", mtest.NextBatch)
		mt.AddMockResponses(firstBatch, killCursors)

		repo := repository.New(mt.Coll)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		results, err := repo.GetAll(ctx, 2, 2)

		assert.NoError(t, err)
		assert.Len(t, results, 2)
		assert.Equal(t, "create user", results[0].Event)
		assert.Equal(t, "update user", results[1].Event)

	})
}
