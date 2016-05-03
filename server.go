package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Get("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Success!")
	})

	fmt.Printf("Running server on port 8080 ...\n")
	e.Run(standard.New(":8080"))
}
