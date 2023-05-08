package view

import (
	"isjhar/template/echo-golang/domain/entities"
	"isjhar/template/echo-golang/utils"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gopkg.in/guregu/null.v4"
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
			raw, err := utils.GetUser(token)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, nil)
			}
			userRaw := raw.(map[string]interface{})
			authorizedContext.User = entities.User{
				ID:       int64(userRaw["id"].(float64)),
				Username: userRaw["username"].(string),
				Name:     null.StringFrom(userRaw["name"].(string)),
			}
			return next(authorizedContext)
		}
	}

}
