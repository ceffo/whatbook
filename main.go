package main

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/tylerb/graceful.v1"
	api "whatbook.com/whatbook/api-library"
	"whatbook.com/whatbook/repository/jsonfile"
	whatbook "whatbook.com/whatbook/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

const (
	port            = 5000
	shutdownTimeout = 10 * time.Second
	dbFile          = "./books/db.json"
)

func main() {
	e, err := setupHandlers()

	if err != nil {
		fmt.Printf("Failed to setup handlers: %s\n", err.Error())
		os.Exit(-1)
	}

	runServer(e, port, shutdownTimeout)
}

func setupHandlers() (*echo.Echo, error) {
	db, err := jsonfile.NewDb(dbFile)
	if err != nil {
		return nil, err
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.INFO)
	e.Logger.SetOutput(os.Stdout)

	svc := whatbook.NewService(db, e.Logger)
	api.RegisterHandlers(e, svc)

	return e, nil
}

func runServer(e *echo.Echo, port int, timeout time.Duration) {
	addr := fmt.Sprintf(":%d", port)
	e.Server.Addr = addr
	e.Logger.Infof("WhatBook: Listening on %s...", addr)
	e.Logger.Info(graceful.ListenAndServe(e.Server, timeout))
}
