package auth

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func registerRoutes(e *echo.Echo, userService *usecase.Service) {
	e.POST("/auth/login", func(c echo.Context) error {
		var credentials LoginRequest
		if err := c.Bind(&credentials); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":   "Bad Request",
				"details": "Required credentials email or password",
			})
		}

		if !isEmail(credentials.Email) || !isPassword(credentials.Password) {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":   "Bad Request",
				"details": "Invalid credentials",
			})
		}

		user, err := userService.FindByEmail(credentials.Email)
		if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error":   "UnAuthorized",
				"details": "Wrong credentials",
			})
		}

		token, err := generateToken(user.ID, user.Name)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error":   "Internal Server Error",
				"details": "Something go wrong",
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"user_id": user.ID,
			"_token":  token,
		})
	})
}

func Run(e *echo.Echo) {
	userService := users.NewService()
	registerRoutes(e, userService)
}
