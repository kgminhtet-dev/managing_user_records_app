package handler

import (
	"github.com/google/uuid"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"regexp"
)

func HandleUserHandlerError(c echo.Context, err error) error {
	switch err {
	case usecase.ErrEmailAlreadyExist:
		return c.JSON(
			http.StatusConflict,
			map[string]string{
				"error":   "Conflict",
				"details": "Email already exists",
			},
		)
	case usecase.ErrUserNotFound:
		return c.JSON(http.StatusNotFound,
			map[string]string{
				"error":   "Not Found",
				"details": "User not found",
			})
	case usecase.ErrInternal:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Internal Server Error",
				"details": "Something went wrong",
			},
		)
	default:
		return c.JSON(
			http.StatusInternalServerError,
			map[string]string{
				"error":   "Unknown Error",
				"details": "An unexpected error occurred",
			},
		)
	}
}

func NewPayload(userId string, data any) *mqueue.Payload {
	return &mqueue.Payload{
		UserID: userId,
		Data:   data,
	}
}

func IsUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

func IsPassword(password string) bool {
	return len(password) >= 8
}

func IsEmail(email string) bool {
	var emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
