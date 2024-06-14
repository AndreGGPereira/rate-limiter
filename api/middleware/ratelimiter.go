package middleware

import (
	"context"
	"net"
	"net/http"
	"rate-limiter/usecase"
	"strings"

	"github.com/labstack/echo/v4"
)

type RateLimiter struct {
	UseCase *usecase.RateLimiter
}

func NewRateLimiterMiddleware(useCase *usecase.RateLimiter) *RateLimiter {
	return &RateLimiter{
		UseCase: useCase,
	}
}

func (rl *RateLimiter) RaceLimiterMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		ctx := context.Background()

		token := c.Request().Header.Get("API_KEY")
		if token != "" {
			if !rl.UseCase.AllowRequest(ctx, token, true) {
				return echo.NewHTTPError(http.StatusTooManyRequests, "you have reached the maximum number of requests or actions allowed within a certain time frame")
			}
			return next(c)
		} else {
			if !rl.UseCase.AllowRequest(ctx, getIP(c.Request()), false) {
				return echo.NewHTTPError(http.StatusTooManyRequests, "you have reached the maximum number of requests or actions allowed within a certain time frame")
			}
			return next(c)
		}
	}
}

func getIP(r *http.Request) string {

	if forwarded := r.Header.Get("X-Forwarded-For"); forwarded != "" {
		return strings.Split(forwarded, ",")[0]
	} else {
		host, _, err := net.SplitHostPort(r.RemoteAddr)

		if err != nil {
			return ""
		}
		return host
	}
}
