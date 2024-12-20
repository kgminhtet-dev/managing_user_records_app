package data

import (
	"context"
	"testing"
)

func TestNewMongo(t *testing.T) {
	const uri = "mongodb://127.0.0.1:27017"
	client := NewMongo(context.TODO(), uri)
	if client == nil {
		t.Errorf("Mongo client not to be nil")
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			t.Errorf("Error disconnecting from MongoDB: %v", err)
		}
	}()
}
