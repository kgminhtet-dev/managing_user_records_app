package mqueue

import (
	"context"
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
	var sub Subscriber = func(ctx context.Context, msg any) error {
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
	var sub1 Subscriber = func(ctx context.Context, msg any) error {
		message := msg.(*Message)
		_, err := testEnv.writer.Write([]byte("Subscriber 1: " + message.Payload.(string)))
		return err
	}
	broker.receive("UserCreated", sub1)
	broker.deliver("UserCreated", "User data")

	wg.Wait()

	assert.Contains(t, testEnv.writer.String(), "User data")
}
