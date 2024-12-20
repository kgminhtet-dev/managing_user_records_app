package mqueue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReceiveMessage(t *testing.T) {
	broker := newBroker()
	event := "UserCreated"
	payload := "User data"
	message := NewMessage(event, payload)
	broker.receive(message)

	assert.Equal(t, payload, broker.getEvent(event))
}
