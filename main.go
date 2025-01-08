package main

import (
	"isjhar/template/echo-golang/data/repositories"
	"isjhar/template/echo-golang/docs"
	"isjhar/template/echo-golang/utils"
	"isjhar/template/echo-golang/view"
	"isjhar/template/echo-golang/view/routers"
	"log"
	"time"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

// @title           Silog API
// @version         1.0
// @description     Silog API.
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	docs.SwaggerInfo.Title = "Silog API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = utils.GetBaseUrl()
	docs.SwaggerInfo.Schemes = []string{"https", "http"}

	e := echo.New()
	e.HideBanner = true
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	e.HTTPErrorHandler = view.CustomHTTPErrorHandler

	buildMiddleware(e)

	environment := utils.GetEnvironment()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	switch environment {
	case utils.LOCAL, utils.DEVELOPMENT:
	case utils.PROD:
		output := &repositories.ElasticsearchRepository{}
		log.SetOutput(output)
	default:
		output := &utils.CustomLogger{}
		log.SetOutput(output)
	}

	logLevel := utils.GetLogLevel()
	if logLevel == utils.LogInfoLevel {
		utils.LogLevel = utils.LogInfoLevel
	} else if logLevel == utils.LogWarningLevel {
		utils.LogLevel = utils.LogWarningLevel
	} else if logLevel == utils.LogErrorLevel {
		utils.LogLevel = utils.LogErrorLevel
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
