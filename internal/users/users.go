package users

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Run(e *echo.Echo) {
	userRoute := e.Group("/api/v1")

	userRoute.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Response from user")
	})

}
