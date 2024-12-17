package users

type Handler struct {
	service *Service
}

func newHandler(service *Service) *Handler {
	return &Handler{service: service}
}
