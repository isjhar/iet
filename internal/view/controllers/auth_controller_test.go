package controllers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/isjhar/iet/internal/view/dto"
	"github.com/isjhar/iet/utils"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func TestLogin_ReturnSuccess(t *testing.T) {
	body := dto.LoginParams{
		Username: "admin",
		Password: "1234",
	}
	bodyJson, _ := json.Marshal(&body)

	e := echo.New()
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(),
	}
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(bodyJson)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/auth/login")

	h := Login()

	h(c)

	if http.StatusOK != rec.Code {
		bodyBytes, err := io.ReadAll(rec.Body)
		if err != nil {
			log.Fatalln(err)
		}
		t.Fatalf("Call error: %v -> %s", rec.Code, string(bodyBytes))
	}
}
