package handler

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/common"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/event"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type UserClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type Handler struct {
	service *usecase.Service
	mq      *mqueue.Mqueue
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
	users, err := h.service.GetUsers(int(page), limit)
	if err != nil {
		return handleUserHandlerError(c, err)
	}

	claim := c.Get("user").(*jwt.Token).Claims.(*common.UserClaims)
	go h.mq.Publish(
		event.UsersFetched,
		newPayload(claim.ID, map[string]int64{"page": page}),
	)

	return c.JSON(http.StatusOK, map[string]any{
		"data":   users,
		"paging": map[string]int64{"previous": page, "next": page + 2},
	})
}

func (h *Handler) GetUser(c echo.Context) error {
	id := c.Param("id")

	if !isUUID(id) {
		return c.JSON(http.StatusBadRequest, BadRequestResponse("Invalid user id"))
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		return handleUserHandlerError(c, err)
	}

	claim := c.Get("user").(*jwt.Token).Claims.(*common.UserClaims)
	go h.mq.Publish(
		event.UserFetched,
		newPayload(claim.ID, map[string]string{"id": id}),
	)

	return c.JSON(
		http.StatusOK,
		map[string]any{
			"data": user,
		})
}

func (h *Handler) CreateUser(c echo.Context) error {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			},
		)
	}

	if input.Name == "" || !isEmail(input.Email) || !isPassword(input.Password) {
		return c.JSON(
			http.StatusBadRequest,
			BadRequestResponse("Invalid user data"))
	}

	var user data.User
	user.Name = input.Name
	user.Email = input.Email
	user.Password = input.Password

	err := h.service.CreateUser(&user)
	if err != nil {
		return handleUserHandlerError(c, err)
	}

	claim := c.Get("user").(*jwt.Token).Claims.(*common.UserClaims)
	go h.mq.Publish(
		event.UserCreated,
		newPayload(
			claim.ID,
			map[string]string{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			}),
	)

	return c.JSON(http.StatusCreated, nil)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	var user data.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{
				"error":   "Bad Request",
				"details": "Invalid data",
			},
		)
	}

	id := c.Param("id")
	user.ID = id

	if user.Name == "" || !isUUID(user.ID) || !isEmail(user.Email) {
		return c.JSON(
			http.StatusBadRequest,
			BadRequestResponse("Invalid user data"))
	}

	err := h.service.UpdateUser(user.ID, &user)
	if err != nil {
		return handleUserHandlerError(c, err)
	}

	claim := c.Get("user").(*jwt.Token).Claims.(*common.UserClaims)
	go h.mq.Publish(
		event.UserUpdated,
		newPayload(
			claim.ID,
			map[string]string{
				"id":    user.ID,
				"name":  user.Name,
				"email": user.Email,
			}),
	)

	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if !isUUID(id) {
		return c.JSON(http.StatusBadRequest, BadRequestResponse("Invalid user id"))
	}

	err := h.service.DeleteUser(id)
	if err != nil {
		return handleUserHandlerError(c, err)
	}

	claim := c.Get("user").(*jwt.Token).Claims.(*common.UserClaims)
	go h.mq.Publish(
		event.UserDeleted,
		newPayload(claim.ID, map[string]string{"id": id}),
	)

	return c.JSON(http.StatusOK, nil)
}

func New(mq *mqueue.Mqueue, service *usecase.Service) *Handler {
	return &Handler{mq: mq, service: service}
}
