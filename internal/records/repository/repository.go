package repository

import (
	"context"
	"fmt"

	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func (r *Repository) Create(ctx context.Context, record *data.UserRecord) error {
	_, err := r.collection.InsertOne(ctx, record)
	if err != nil {
		return fmt.Errorf("failed to insert record: %w", err)
	}
	return nil
}

func (r *Repository) GetById(ctx context.Context, id string) (*data.UserRecord, error) {
	var record data.UserRecord
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&record)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("record not found with id %s", id)
		}
		return nil, fmt.Errorf("failed to get record: %w", err)
	}
	return &record, nil
}

func (r *Repository) GetAll(ctx context.Context, start, limit int) ([]*data.UserRecord, error) {
	filter := bson.M{}
	opts := options.Find().
		SetSkip(int64(start)).
		SetLimit(int64(limit)).
		SetSort(bson.D{{Key: "_id", Value: 1}})

	cursor, err := r.collection.Find(ctx, filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve records: %w", err)
	}
	defer cursor.Close(ctx)

	var records []*data.UserRecord
	if err := cursor.All(ctx, &records); err != nil {
		return nil, fmt.Errorf("failed to decode records: %w", err)
	}

	if len(records) == 0 {
		return []*data.UserRecord{}, nil
	}

	return records, nil
}

func New(collection *mongo.Collection) *Repository {
	return &Repository{collection: collection}
}
