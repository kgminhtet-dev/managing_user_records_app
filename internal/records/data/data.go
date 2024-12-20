package data

import (
	"context"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func DisconnectDatabase(ctx context.Context, db *mongo.Database) {
	if err := db.Client().Disconnect(ctx); err != nil {
		panic(err)
	}
}

func ConnectDatabase(ctx context.Context, cfg *config.Config) *mongo.Database {
	client := NewMongo(ctx, cfg.Database.Url)
	database := client.Database(cfg.Database.Name)

	var result bson.M
	if err := database.RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
		panic(err)
	}

	log.Println("Successfully connected to MongoDB!")
	return database
}
