package main

import (
	"isjhar/template/echo-golang/utils"
	"isjhar/template/echo-golang/view/entities"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	e.Use(middleware.Recover())

	environment := utils.GetEnvironment()
	if environment != utils.DEVELOPMENT {
		output := &utils.CustomLogger{}
		e.Logger.SetOutput(output)
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: output,
		}))

		e.Logger.SetLevel(4)
	}

	//CORS Config
	CORSConfig := middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}
	e.Use(middleware.CORSWithConfig(CORSConfig))

	e.Static("/docs", "docs")

	public := e.Group("")

	public.GET("/health", health)
	public.OPTIONS("/health", health)

	private := e.Group("")
	private.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(utils.GetJwtSecret()),
		SigningMethod: "HS512",
	}))

	e.Logger.Fatal(e.Start(":1323"))
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, entities.ApiResponse{
		Message: "still alive",
		Data:    nil,
	})
}
