package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Home(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to RESTFULL API myProperty Kampus Merdeka")
}
