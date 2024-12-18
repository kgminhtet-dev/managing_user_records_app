package users

import "github.com/labstack/echo/v4"

type Handler struct {
	service *Service
}

func (h *Handler) GetUsers(c echo.Context) error {
	return nil
}

func (h *Handler) GetUser(c echo.Context) error {
	return nil
}

func (h *Handler) CreateUser(c echo.Context) error {
	return nil
}

func (h *Handler) UpdateUser(c echo.Context) error {
	return nil
}

func (h *Handler) DeleteUser(c echo.Context) error {
	return nil
}

func newHandler(service *Service) *Handler {
	return &Handler{service: service}
}
