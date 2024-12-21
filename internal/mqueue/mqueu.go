package mqueue

import "sync"

type Mqueue struct {
	broker *broker
}

func (m *Mqueue) Publish(event string, payload any) {
	m.broker.deliver(event, payload)
}

func (m *Mqueue) Subscribe(event string, s Subscriber) {
	m.broker.receive(event, s)
}

func New(wg *sync.WaitGroup) *Mqueue {
	return &Mqueue{broker: newBroker(wg)}
}
