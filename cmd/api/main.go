package main

import (
	"context"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/auth"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/mqueue"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/records"
	"github.com/kgminhtet-dev/managing_user_records_app/internal/users"
	"github.com/labstack/echo/v4"
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
	e := echo.New()
	var wg sync.WaitGroup
	mq := mqueue.New(&wg)

	users.Run(mq, e)
	records.Run(mq, e)
	auth.Run(e)

	wg.Add(1)
	go func() {
		defer wg.Done()
		e.Logger.Fatal(e.Start(os.Getenv("HOST") + ":" + os.Getenv("PORT")))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		e.Logger.Info("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}()

	wg.Wait()
	log.Println("Application exited...")
}
