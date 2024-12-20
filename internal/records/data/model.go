package data

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRecord struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Event     string             `json:"event" bson:"event"`
	Data      any                `json:"data" bson:"data"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}
