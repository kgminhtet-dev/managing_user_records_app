package handler

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service *usecase.Service
}

func (h *Handler) GetUsers(c echo.Context) error {
	pageParam := c.QueryParam("page")
	page, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error":   "Bad request",
				"details": "Invalid page number",
			})
	}

	limit := 10
	user, err := h.service.GetUsers(int(page), limit)
	switch err {
	case usecase.ErrInternal:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			},
		)
	default:
		return c.JSON(
			http.StatusOK,
			map[string]any{
				"data":   user,
				"paging": map[string]int64{"previous": page, "next": page + 2},
			},
		)
	}
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