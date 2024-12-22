package records

import (
	"context"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/config"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/data"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/handler"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/repository"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records/usecase"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

func Run(q *mqueue.Mqueue, r *echo.Group) {
	cfg := config.LoadConfig(os.Getenv("RECORD_CONFIG_PATH"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	database := data.ConnectDatabase(ctx, cfg)
	collection := database.Collection(cfg.Database.Collection)
	repo := repository.New(collection)
	service := usecase.NewService(repo)
	h := handler.New(service)
	handler.RegisterRoutes(r, h)
	handler.RegisterSubscribers(q, h)
}
