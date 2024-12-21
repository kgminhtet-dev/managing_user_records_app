package mqueue

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
)

type TestEnvironment struct {
	writer *strings.Builder
}

func TestReceiveMessage(t *testing.T) {
	var wg sync.WaitGroup

	broker := newBroker(&wg)
	var sub Subscriber = func(payload any) error {
		return nil
	}
	broker.receive("UserCreated", sub)

	assert.Len(t, broker.messages, 1)
	assert.NotNil(t, broker.messages["UserCreated"])

	wg.Wait()
}

func TestDeliveredMessage(t *testing.T) {
	testEnv := &TestEnvironment{
		writer: &strings.Builder{},
	}

	var wg sync.WaitGroup
	broker := newBroker(&wg)
	var sub1 Subscriber = func(payload any) error {
		_, err := testEnv.writer.Write([]byte("Subscriber 1: " + payload.(string)))
		return err
	}
	broker.receive("UserCreated", sub1)
	broker.deliver("UserCreated", "User data")

	wg.Wait()

	t.Log("Data ", testEnv.writer.String())
	assert.Contains(t, testEnv.writer.String(), "User data")
}
