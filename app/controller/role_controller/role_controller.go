package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/role_service"

	"github.com/labstack/echo/v4"
)

type RolesController struct {
	roleService *service.Roleservice
}

func NewRolesController(es *service.Roleservice) *RolesController {
	return &RolesController{
		roleService: es,
	}
}

func (s *RolesController) NewRoles(ctx echo.Context) error {
	payload := new(model.Roles)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.roleService.NewRoles(payload); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success create record",
	})
}

func (s *RolesController) GetRolesByID(ctx echo.Context) error {
	nip := ctx.Param("id")

	Roles, err := s.roleService.GetRolesByID(nip)
	if err != nil {
		if err.Error() == "there is no record with that id" {
			return ctx.JSON(http.StatusNotFound, model.Response{
				Status:  "failed",
				Message: err.Error(),
			})
		}

		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    Roles,
	})
}

func (s *RolesController) GetRoles(ctx echo.Context) error {
	Roles, err := s.roleService.GetRoles()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    Roles,
	})
}

func (s *RolesController) UpdateRoles(ctx echo.Context) error {
	payload := new(model.Roles)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.roleService.UpdateRoles(payload); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success update record",
	})
}
