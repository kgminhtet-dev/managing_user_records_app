package data

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type UserRecord struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Event     string             `json:"event" bson:"event"`
	Data      *data.User         `json:"data" bson:"data"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"update_at" bson:"updated_at"`
}
