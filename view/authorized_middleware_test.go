package view

import (
	"isjhar/template/echo-golang/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestAuthorizedUser(t *testing.T) {

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	h := AuthorizedUser("header")(func(c echo.Context) error {
		return c.String(http.StatusOK, "test")
	})
	userPayload := make(map[string]interface{})
	userPayload["id"] = 1
	userPayload["username"] = "admin"
	token, err := utils.GenerateJWT(userPayload)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set(echo.HeaderAuthorization, "bearer "+token)

	h(c)

	if http.StatusOK != rec.Code {
		t.Fatalf("Call error")
	}
}
