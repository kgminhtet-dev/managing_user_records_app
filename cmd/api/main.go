package main

import (
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users"
	"github.com/labstack/echo/v4"
	"github.com/mr-kmh/envify"
	"net/http"
	"os"
)

func init() {
	envify.Load()
}

func main() {
	e := echo.New()

	e.GET("/api/v1", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Managing User Record API!!!")
	})
	users.Run(e)

	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
