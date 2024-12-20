package usecase

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
