package main

import (
	"github.com/labstack/echo/v4"
	"github.com/mr-kmh/envify"
	"net/http"
	"os"
)

func main() {
	e := echo.New()
	envify.Load()

	e.GET("/api/v1", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Managing User Record API!!!")
	})

	e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
}
