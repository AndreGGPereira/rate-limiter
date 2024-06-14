package handlers

import (
	"rate-limiter/usecase"

	"rate-limiter/api/controller"

	"github.com/labstack/echo/v4"
)

// MakeDriverHandlers make url handlers
func MakeHomeHandlers(e *echo.Echo, usecase *usecase.RateLimiter) {

	c := controller.NewControler()

	e.GET("", c.Home)
}
