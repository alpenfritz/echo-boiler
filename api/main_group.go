package api

import (
	"echo-boiler/handlers"

	"github.com/labstack/echo/v4"
)

// MainGroup - Routes for main group
func MainGroup(e *echo.Echo) {
	e.GET("/health-check", handlers.HealthCheck)
}
