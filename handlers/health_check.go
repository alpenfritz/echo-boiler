package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-boiler/renderings"
)

// HealthCheck - Health Check Handler
func HealthCheck(c echo.Context) error {
	c.Logger().Debugf("RequestID: %s", c.Response().Header().Get(echo.HeaderXRequestID))
	resp := renderings.HealthCheckResponse{
		Message: "Everything is good!",
	}
	return c.JSON(http.StatusOK, resp)
}
