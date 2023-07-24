package mdl

import (
	"net/http"
	"presensi/helper"
	"presensi/model"

	"github.com/labstack/echo/v4"
)

type Middleware struct {
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}

func (middleware *Middleware) OnlyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("access-token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, model.Response{
				Status:  "failed",
				Message: "missing access token",
			})
		}

		token := cookie.Value

		role, ok := helper.VerifyJWT(token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, model.Response{
				Status:  "failed",
				Message: "unauthorized access",
			})
		}

		if role != "ADMIN" {
			return c.JSON(http.StatusUnauthorized, model.Response{
				Status:  "failed",
				Message: "unauthorized access",
			})
		}

		return next(c)

	}
}

func (middleware *Middleware) AllRole(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		cookie, err := c.Cookie("access_token")
		if err != nil {
			return c.JSON(http.StatusUnauthorized, model.Response{
				Status:  "failed",
				Message: "missing access token",
			})
		}

		token := cookie.Value

		_, ok := helper.VerifyJWT(token)
		if !ok {
			return c.JSON(http.StatusUnauthorized, model.Response{
				Status:  "failed",
				Message: "unauthorized access",
			})
		}

		return next(c)

	}
}
