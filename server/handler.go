package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func landing(c echo.Context) error {
	return c.String(http.StatusOK, "Api Rodando")
}
