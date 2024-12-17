package main

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	cfg := config.Load()
	e := echo.New()

	e.GET("/api/v1", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Managing User Record API!!!")
	})

	e.Logger.Fatal(e.Start(cfg.Host + ":" + cfg.Port))
}