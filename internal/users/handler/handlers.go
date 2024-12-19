package handler

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
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
	id := c.Param("id")
	if !isUUID(id) {
		return c.JSON(http.StatusBadRequest, BadRequestResponse("Invalid user id"))
	}

	user, err := h.service.GetUserById(id)
	switch err {
	case usecase.ErrInternal:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			},
		)
	case usecase.ErrUserNotFound:
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error":   "Not Found",
				"details": "User not found",
			},
		)
	}

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
	switch err {
	case usecase.ErrEmailAlreadyExist:
		return c.JSON(
			http.StatusConflict,
			map[string]string{
				"error":   "Conflict",
				"details": "Email already exists",
			},
		)
	case usecase.ErrInternal:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			},
		)
	}

	return c.JSON(http.StatusCreated, nil)
}

func (h *Handler) UpdateUser(c echo.Context) error {
	var user data.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
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
	switch err {
	case usecase.ErrUserNotFound:
		return c.JSON(
			http.StatusNotFound,
			map[string]string{
				"error":   "Not Found",
				"details": "User not found",
			})
	case usecase.ErrInternal:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			},
		)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if !isUUID(id) {
		return c.JSON(http.StatusBadRequest, BadRequestResponse("Invalid user id"))
	}

	err := h.service.DeleteUser(id)
	switch err {
	case usecase.ErrInternal:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			},
		)
	}

	return c.JSON(http.StatusOK, nil)
}

func New(service *usecase.Service) *Handler {
	return &Handler{service: service}
}
