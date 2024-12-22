package main

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/auth"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/common"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mr-kmh/envify"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func init() {
	envify.Load()
}

func main() {
	var wg sync.WaitGroup

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	auth.Run(e)

	r := e.Group("/api/v1")
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(common.UserClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET_TOKEN")),
	}
	r.Use(echojwt.WithConfig(config))
	mq := mqueue.New(&wg)
	users.Run(mq, r)
	records.Run(mq, r)

	wg.Add(1)
	go func() {
		defer wg.Done()
		e.Logger.Fatal(e.Start(common.GetURI()))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	wg.Add(1)
	go func() {
		<-quit
		e.Logger.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer func() {
			cancel()
			wg.Done()
		}()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	wg.Wait()
	log.Println("Application exited...")
}
