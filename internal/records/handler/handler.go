package handler

import (
	"context"
	"fmt"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/event"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Handler struct {
	service *usecase.Service
}

func RegisterRoutes(router *echo.Group, handler *Handler) {
	router.GET("/records", handler.GetRecords)
}

func RegisterSubscribers(q *mqueue.Mqueue, handler *Handler) {
	q.Subscribe(event.UserFetched, handler.CrateRecordSubscriber)
	q.Subscribe(event.UsersFetched, handler.CrateRecordSubscriber)
	q.Subscribe(event.UserCreated, handler.CrateRecordSubscriber)
	q.Subscribe(event.UserUpdated, handler.CrateRecordSubscriber)
	q.Subscribe(event.UserDeleted, handler.CrateRecordSubscriber)
}

func (h *Handler) CrateRecordSubscriber(ctx context.Context, message any) error {
	msg, ok := message.(*mqueue.Message)
	if !ok {
		return fmt.Errorf("invalid message format")
	}

	payload, ok := msg.Payload.(*mqueue.Payload)
	if !ok || !payload.Validate() {
		return fmt.Errorf("invalid payload format")
	}

	return h.service.CreateRecord(ctx, msg.Event, payload)
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

	var response any
	var statusCode int
	switch err {
	case usecase.ErrDatabaseError:
		statusCode = http.StatusInternalServerError
		response = map[string]string{
			"error":   "Internal Server Error",
			"details": "Something went wrong",
		}
	case usecase.ErrNotFound:
		statusCode = http.StatusOK
		response = map[string]interface{}{
			"data": nil,
		}
	default:
		statusCode = http.StatusOK
		response = map[string]interface{}{
			"data": records,
		}
	}

	return c.JSON(statusCode, response)
}

func New(service *usecase.Service) *Handler {
	return &Handler{service: service}
}
