package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// V1Main - Handler for v1 main
func V1Main(c echo.Context) error {
	return c.String(http.StatusOK, "You are on the v1 main page")
}
