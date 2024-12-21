package mqueue

import (
	"log"
	"sync"
)

type Subscriber func(payload any) error
type Subscribers []Subscriber

type broker struct {
	mu       sync.RWMutex
	wg       *sync.WaitGroup
	messages map[string]Subscribers
}

func (b *broker) receive(event string, s Subscriber) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if _, ok := b.messages[event]; ok {
		b.messages[event] = append(b.messages[event], s)
	} else {
		b.messages[event] = Subscribers{s}
	}
}

func (b *broker) deliver(event string, payload any) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if fns, ok := b.messages[event]; ok {
		for _, fn := range fns {
			b.wg.Add(1)
			go (func() {
				defer b.wg.Done()
				if err := fn(payload); err != nil {
					log.Printf("event: %s, error: %v", event, err)
				}
			})()
		}
	}
}

func newBroker(wg *sync.WaitGroup) *broker {
	subscribers := make(map[string]Subscribers)
	return &broker{messages: subscribers, wg: wg}
}
