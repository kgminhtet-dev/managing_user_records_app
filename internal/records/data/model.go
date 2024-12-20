package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRecord struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Event     string             `json:"event" bson:"event"`
	Data      any                `json:"data" bson:"data"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"update_at" bson:"updated_at"`
}
