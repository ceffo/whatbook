package main

import (
	"fmt"

	api "whatbook.com/whatbook/api-library"

	impl "whatbook.com/whatbook/impl"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Printf("WhatBook\n")

	setupHandlers()
}

func setupHandlers() {
	var myAPI impl.WhatBook
	e := echo.New()
	e.Use(middleware.Logger())

	api.RegisterHandlers(e, &myAPI)

	e.Start(":5000")
}
