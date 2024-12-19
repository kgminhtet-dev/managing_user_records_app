package testutils

import "testing"

func TestSetup(t *testing.T) {
	db := Setup()
	if db == nil {
		t.Fatalf("Expected database not to be nil")
	}
}
