package main

import (
	"blitztracker_api/src/config"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	// load config
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	e := echo.New()

	e.Logger.SetLevel(log.INFO)

	//use cors config
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		time.Sleep(2 * time.Second)
		return c.JSON(http.StatusOK, "Welcome to Toothless Tarantula BlitzTracker API!")
	})

	//all function handlers goes here
	//init db
	// DB := db.Init()

	// Start server
	go func() {
		if err := e.Start(":" + env.PORT); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
