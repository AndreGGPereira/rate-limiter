package server

import (
	"context"
	"log"
	"net/http"
	"rate-limiter/api/handlers"
	"rate-limiter/api/middleware"
	"rate-limiter/config"
	"rate-limiter/infra/cache"
	"rate-limiter/usecase"

	"github.com/labstack/echo/v4"
)

func Execute() {

	e := echo.New()
	ctx := context.Background()

	cache, err := cache.NewConnection(
		ctx,
		config.Cache().Host,
		config.Cache().Port,
		config.Cache().Pwd,
		config.Cache().DB,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.Close()

	usecase := usecase.NewRateLimiter(cache)

	middle := middleware.NewRateLimiterMiddleware(usecase)

	e.Use(middle.RaceLimiterMiddleware)

	handlers.MakeHomeHandlers(e, usecase)

	if err := e.Start(":" + config.Server().Port); err != http.ErrServerClosed {
		log.Fatal(err)

	}
}
