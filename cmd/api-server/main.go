package main

import (
	"log"
	"time"

	"github.com/isjhar/iet/internal/config"
	"github.com/isjhar/iet/internal/data/repositories"
	"github.com/isjhar/iet/internal/docs"
	"github.com/isjhar/iet/internal/view"
	"github.com/isjhar/iet/internal/view/routers"
	"github.com/isjhar/iet/pkg"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

// @title           Service API
// @version         1.0
// @description     Service API.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	config.LoadConfig()

	docs.SwaggerInfo.Title = config.Swagger.Title.String
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = pkg.GetBaseUrl()
	docs.SwaggerInfo.Schemes = []string{config.Swagger.Scheme.String}

	e := echo.New()
	e.HideBanner = true
	e.Validator = &pkg.CustomValidator{
		Validator: validator.New(),
	}
	e.HTTPErrorHandler = view.CustomHTTPErrorHandler

	buildMiddleware(e)

	environment := pkg.GetEnvironment()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	switch environment {
	case pkg.LOCAL, pkg.DEVELOPMENT:
	case pkg.PROD:
		output := &repositories.ElasticsearchRepository{}
		log.SetOutput(output)
	default:
		output := &pkg.CustomLogger{}
		log.SetOutput(output)
	}

	logLevel := pkg.GetLogLevel()
	if logLevel == pkg.LogInfoLevel {
		pkg.LogLevel = pkg.LogInfoLevel
	} else if logLevel == pkg.LogWarningLevel {
		pkg.LogLevel = pkg.LogWarningLevel
	} else if logLevel == pkg.LogErrorLevel {
		pkg.LogLevel = pkg.LogErrorLevel
	}

	routers.Route(e)

	e.Logger.Fatal(e.Start(":1323"))
}

func buildMiddleware(e *echo.Echo) {
	//CORS Config
	CORSConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}
	e.Use(middleware.CORSWithConfig(CORSConfig))

	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(50), Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
	}))
	e.Use(view.ElapsedTimeMiddleware())

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 1 * time.Minute,
		Skipper: func(c echo.Context) bool {
			return c.Request().Header.Get("Upgrade") == "websocket"
		},
	}))
}
