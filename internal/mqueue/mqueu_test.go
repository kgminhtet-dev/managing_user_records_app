package mqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMqueue_Publish(t *testing.T) {
	mqueue := New()
	event := "UserCreated"
	payload := "User data"
	mqueue.Publish(event, payload)
	storedPayload := mqueue.broker.getEvent(event)

	assert.Equal(t, payload, storedPayload)
}
