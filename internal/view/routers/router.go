package routers

import (
	"net/http"

	"github.com/isjhar/iet/internal/data/repositories"
	"github.com/isjhar/iet/internal/view"
	"github.com/isjhar/iet/internal/view/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo) {
	public := e.Group("")
	public.GET("/health", health)
	public.OPTIONS("/health", health)

	AuthRouter(public)

	jwtRepository := repositories.JwtRepository{}

	private := e.Group("")
	private.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(jwtRepository.GetJwtSecret()),
		SigningMethod: "HS512",
	}))
	private.Use(view.AuthorizedUser("header"))
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.ApiResponse{
		Message: "still alive",
	})
}
