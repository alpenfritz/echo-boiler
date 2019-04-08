package api

import (
	"echo-boiler/handlers"

	"github.com/labstack/echo/v4"
)

// V1Group - Routes for v1 group
func V1Group(g *echo.Group) {
	g.GET("/main", handlers.V1Main)
}
