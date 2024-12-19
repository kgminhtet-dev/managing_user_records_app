package users

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/config"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/handler"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/repository"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func routes(router *echo.Group, handlers *handler.Handler) {
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
}

func Run(e *echo.Echo) {
	cfg, err := config.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	db := data.New(cfg.Database)
	repo := repository.New(db)
	service := usecase.NewService(repo)
	h := handler.New(service)

	userRoute := e.Group("api/v1")
	userRoute.GET("users", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Response from user")
	})
	routes(userRoute, h)
}
