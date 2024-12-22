package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetURI(t *testing.T) {
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "8080")
	uri := GetURI()
	assert.Equal(t, "localhost:8080", uri)
}

func TestGenerateLink(t *testing.T) {
	uri := "http://localhost:8080"
	endpoint := "users"
	id := "123"
	expected := map[string]map[string]string{
		"self": {
			"href":   "http://localhost:8080/users/123",
			"method": "GET",
		},
		"collection": {
			"href":   "http://localhost:8080/users",
			"method": "GET",
		},
		"next": {
			"href":   "http://localhost:8080/users?page=3",
			"method": "GET",
		},
		"prev": {
			"href":   "http://localhost:8080/users?page=1",
			"method": "GET",
		},
	}
	result := GenerateLinks(uri, endpoint, id, 2)
	assert.Equal(t, expected, result)
}
