package auth

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users"
	userHandler "github.com/kgminhtet-dev/managing_user_records_app/internal/users/handler"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func registerRoutes(g *echo.Group, userService *usecase.Service) {
	g.POST("/auth/login", func(c echo.Context) error {
		var credentials LoginRequest
		if err := c.Bind(&credentials); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error":   "Bad Request",
				"details": "Required credentials email or password",
			})
		}

		if !userHandler.IsEmail(credentials.Email) || !userHandler.IsPassword(credentials.Password) {
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

		return c.JSON(http.StatusOK, nil)
	})
}

func Run(e *echo.Echo) {
	api := e.Group("/api/v1")
	userService := users.NewService()
	registerRoutes(api, userService)
}
