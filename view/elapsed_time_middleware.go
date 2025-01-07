package view

import (
	"isjhar/template/echo-golang/data/repositories"
	"isjhar/template/echo-golang/utils"
	"time"

	"github.com/labstack/echo/v4"
)

var elasticSearch = repositories.ElasticsearchRepository{}

// ElapsedTimeMiddleware calculates the elapsed time for each request
func ElapsedTimeMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("Upgrade") == "websocket" {
				return next(c)
			}

			start := time.Now()

			// Process request
			err := next(c)

			// Calculate elapsed time
			elapsed := time.Since(start)

			utils.LogInfo("url=%s status_code=%d duration=%d", c.Request().URL.Path, c.Response().Status, elapsed.Milliseconds())

			elasticSearch.LogApi(repositories.LogApiParams{
				Path:       c.Request().URL.Path,
				Method:     c.Request().Method,
				StatusCode: c.Response().Status,
				Duration:   elapsed.Milliseconds(),
				UserAgent:  c.Request().Header.Get("User-Agent"),
			})

			return err
		}
	}
}
