package routers

import (
	"net/http"

	"github.com/isjhar/iet/internal/data/repositories"
	"github.com/isjhar/iet/internal/view"
	"github.com/isjhar/iet/internal/view/dto"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func Route(e *echo.Echo) {
	publicRoute(e)
	privateRoute(e)
}

func publicRoute(e *echo.Echo) {
	public := e.Group("")
	public.GET("/health", health)
	public.OPTIONS("/health", health)

	e.GET("/docs*", echoSwagger.WrapHandler)

	AuthRouter(public)
}

// health godoc
// @Summary      Check server health
// @Tags         Health
// @Produce      json
// @Success      200  {object}  dto.ApiResponse
// @Router       /health [get]
func health(c echo.Context) error {
	return c.JSON(http.StatusOK, dto.ApiResponse{
		Message: "still alive",
	})
}

func privateRoute(e *echo.Echo) {
	jwtRepository := repositories.JwtRepository{}

	private := e.Group("")
	private.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    []byte(jwtRepository.GetJwtSecret()),
		SigningMethod: "HS512",
	}))
	private.Use(view.AuthorizedUser("header"))
}
