package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/api/v1", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, "Managing User Record API!!!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
