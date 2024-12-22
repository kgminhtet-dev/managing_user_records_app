package users

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/config"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/handler"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/repository"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users/usecase"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

var (
	service *usecase.Service
)

func registerRoutes(router *echo.Group, handlers *handler.Handler) {
	router.GET("/users", handlers.GetUsers)
	router.GET("/users/:id", handlers.GetUser)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users/:id", handlers.UpdateUser)
	router.DELETE("/users/:id", handlers.DeleteUser)
}

func readConfig() *config.Config {
	file, err := os.Open(os.Getenv("USER_CONFIG_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Load(file)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

func NewService() *usecase.Service {
	if service != nil {
		return service
	}

	cfg := readConfig()
	db := data.New(cfg.Database)
	repo := repository.New(db)
	service = usecase.NewService(repo)
	return service
}

func Run(q *mqueue.Mqueue, r *echo.Group) {
	h := handler.New(q, NewService())
	registerRoutes(r, h)
}
