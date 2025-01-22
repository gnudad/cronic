package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewServer(cronic *Cronic) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Cronic Scheduler")
	})

	return e
}
