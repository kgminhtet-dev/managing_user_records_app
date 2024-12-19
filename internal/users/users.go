package users

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/config"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/handler"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/repository"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func routes(router *echo.Group, handlers *handler.Handler) {
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

	cfg, err := config.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	db := data.New(cfg.Database)
	repo := repository.New(db)
	service := usecase.NewService(repo)
	handler := handler.New(service)
	routes(userRoute, handler)
}
