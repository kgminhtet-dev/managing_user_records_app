package data

import "testing"

func TestNewMongo(t *testing.T) {
	client := NewMongo()
	if client == nil {
		t.Errorf("Mongo client not to be nil")
	}
}
