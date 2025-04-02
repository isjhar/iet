package view

import (
	"net/http"
	"strings"

	"github.com/isjhar/iet/internal/data/repositories"
	"github.com/isjhar/iet/internal/domain/entities"
	"github.com/isjhar/iet/internal/domain/usecases"

	"github.com/labstack/echo/v4"
)

type AuthorizedContext struct {
	echo.Context
	User entities.User
}

func (c *AuthorizedContext) GetToken(tokenLookup string) string {
	tokenString := ""
	if tokenLookup == "" || tokenLookup == "header" {
		bearerToken := c.Request().Header.Get("Authorization")
		tokenString = strings.Split(bearerToken, " ")[1]
	} else if tokenLookup == "query" {
		tokenString = c.QueryParam("token")
	}
	return tokenString
}

func AuthorizedUser(tokenLookup string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizedContext := &AuthorizedContext{
				Context: c,
			}
			token := authorizedContext.GetToken(tokenLookup)
			getTokenUserUseCase := usecases.GetTokenUserUseCase{
				JwtRepository: repositories.JwtRepository{},
			}
			user, err := getTokenUserUseCase.Execute(c.Request().Context(), token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, nil)
			}
			authorizedContext.User = user
			return next(authorizedContext)
		}
	}

}
