package router

import (
	"echo-boiler/api"
	"echo-boiler/bindings"
	"echo-boiler/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

//New Create a new instance of Echo
func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Validator = new(bindings.Validator)

	// Create groups
	v1 := e.Group("/v1")

	// Set middlewares
	middlewares.SetMainMiddlewares(e)

	// Set main routes
	api.MainGroup(e)

	// Set group routes
	api.V1Group(v1)

	return e
}
