package handler

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service *usecase.Service
}

func SetupRoutes(router *echo.Group, handlers *Handler) {
	router.GET("/records", handlers.GetRecords)
}

func (h *Handler) GetRecords(c echo.Context) error {
	page := c.QueryParam("page")
	pageNo, err := strconv.ParseInt(page, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "Bad request",
			"details": "Invalid page number",
		})
	}

	records, err := h.service.GetRecords(c.Request().Context(), int(pageNo))
	switch err {
	case usecase.ErrDatabaseError:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			},
		)
	case usecase.ErrNotFound:
		return c.JSON(
			http.StatusOK,
			map[string]any{
				"data": nil,
			},
		)
	default:
		return c.JSON(
			http.StatusOK,
			map[string]any{
				"data":   records,
				"paging": map[string]int64{"previous": pageNo, "next": pageNo + 1},
			},
		)
	}
}

func New(service *usecase.Service) *Handler {
	return &Handler{service: service}
}
