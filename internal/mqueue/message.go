package mqueue

type Message struct {
	Event   string
	Payload any
}

type Payload struct {
	UserID string
	Data   any
}

func (p *Payload) Validate() bool {
	if p.UserID == "" || p.Data == nil {
		return false
	}
	return true
}

func NewMessage(event string, payload any) *Message {
	return &Message{
		Event:   event,
		Payload: payload,
	}
}
