package mqueue

import (
	"context"
	"log"
	"sync"
	"time"
)

type Subscriber func(ctx context.Context, msg any) error
type Subscribers []Subscriber

type broker struct {
	mu       sync.RWMutex
	wg       *sync.WaitGroup
	messages map[string]Subscribers
}

func (b *broker) receive(event string, s Subscriber) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.messages[event] = append(b.messages[event], s)
}

func (b *broker) deliver(event string, payload any) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	message := NewMessage(event, payload)
	if fns, ok := b.messages[event]; ok {
		for _, fn := range fns {
			b.wg.Add(1)
			go (func(fn Subscriber) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer func() {
					cancel()
					b.wg.Done()
				}()
				if err := fn(ctx, message); err != nil {
					log.Printf("[mqueue] event: %s, error: %v", event, err)
				}
			})(fn)
		}
	}
}

func newBroker(wg *sync.WaitGroup) *broker {
	subscribers := make(map[string]Subscribers)
	return &broker{messages: subscribers, wg: wg}
}
