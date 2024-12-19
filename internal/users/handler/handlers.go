package handler

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	service *usecase.Service
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

func New(service *usecase.Service) *Handler {
	return &Handler{service: service}
}
