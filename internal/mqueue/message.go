package mqueue

type Message struct {
	Event   string
	Payload any
}

func NewMessage(event string, payload any) *Message {
	return &Message{
		Event:   event,
		Payload: payload,
	}
}
