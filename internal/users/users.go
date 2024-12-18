package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func routes(router *echo.Group, handlers *Handler) {
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
}

func Run(e *echo.Echo) {
	userRoute := e.Group("/api/v1")
	userRoute.GET("/users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Response from user")
	})

	cfg, err := LoadConfig()
	if err != nil {
		e.Logger.Fatal(err)
	}

	db := newDatabase(&cfg.Database)
	repo := newRepository(db)
	service := newService(repo)
	handler := newHandler(service)
	routes(userRoute, handler)
}
