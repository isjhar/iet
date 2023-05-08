package routers

import (
	"isjhar/template/echo-golang/utils"
	"isjhar/template/echo-golang/view"
	"isjhar/template/echo-golang/view/dto"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo) {
	public := e.Group("")
	public.GET("/health", health)
	public.OPTIONS("/health", health)

	AuthRouter(public)

	private := e.Group("")
	private.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(utils.GetJwtSecret()),
		SigningMethod: "HS512",
	}))
	private.Use(view.AuthorizedUser("header"))
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.ApiResponse{
		Message: "still alive",
	})
}
