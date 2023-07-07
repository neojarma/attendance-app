package controller

import (
	"net/http"
	"presensi/helper"
	"presensi/model"
	service "presensi/service/auth_service"
	"time"

	"github.com/labstack/echo/v4"
)

var EXPIRES = time.Now().Add(time.Hour * 24)

type AuthController struct {
	service *service.AuthService
}

func NewAuthController(s *service.AuthService) *AuthController {
	return &AuthController{
		service: s,
	}
}

func (c *AuthController) Auth(ctx echo.Context) error {
	payload := new(model.LoginPayload)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	employee, err := c.service.Auth(payload)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	token, err := helper.GenerateJWT(employee, EXPIRES)
	if err != nil {
		return err
	}

	cookie := helper.GenerateCookie("access-token", token, EXPIRES)
	ctx.SetCookie(cookie)

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "login success",
	})
}
