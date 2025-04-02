package routers

import (
	"github.com/isjhar/iet/internal/view/controllers"

	"github.com/labstack/echo/v4"
)

func AuthRouter(api *echo.Group) {
	api.OPTIONS("/auth/login", controllers.Login())
	api.POST("/auth/login", controllers.Login())
}
