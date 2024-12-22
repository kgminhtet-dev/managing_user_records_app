package testutil

import (
	"context"
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/config"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"time"
)

const collection = "records"

func setupEnvironment() {
	if err := os.Setenv("ENV", "testing"); err != nil {
		log.Fatal(err)
	}
}

func GenerateRandomRecords(size int) []any {
	events := []string{"UserCreated", "UserUpdated", "UserDeleted", "UserFetched"}
	records := make([]any, size)

	for i := range records {
		record := &data.UserRecord{}
		record.ID = primitive.NewObjectID()
		record.UserID = uuid.New().String()
		record.Event = events[i%len(events)]
		record.Data = nil
		record.Timestamp = time.Now()
		records[i] = record
	}

	return records
}

func SeedDatabase(collection *mongo.Collection, records []any) {
	if _, err := collection.InsertMany(context.TODO(), records); err != nil {
		log.Fatal("Error seeding database:", err)
	}
}

func NewEnvironment() (*mongo.Database, *mongo.Collection) {
	setupEnvironment()
	cfg := config.LoadConfig("..")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	database := data.ConnectDatabase(ctx, cfg)
	return database, database.Collection(collection)
}

func Clear(ctx context.Context, database *mongo.Database) {
	if err := database.Collection(collection).Drop(ctx); err != nil {
		panic(err)
	}
}
