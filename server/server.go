package server

import (
	"github.com/labstack/echo/v4"
)

func Start() {
	e := echo.New()
	e.GET("/", landing)
	e.Logger.Fatal(e.Start(":2121"))
}
