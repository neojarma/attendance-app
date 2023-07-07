package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/presence_status_service"

	"github.com/labstack/echo/v4"
)

type PresenceStatusController struct {
	PresenceStatuservice *service.PresenceStatuservice
}

func NewPresenceStatusController(es *service.PresenceStatuservice) *PresenceStatusController {
	return &PresenceStatusController{
		PresenceStatuservice: es,
	}
}

func (s *PresenceStatusController) NewPresenceStatus(ctx echo.Context) error {
	payload := new(model.PresenceStatus)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.PresenceStatuservice.NewPresenceStatus(payload); err != nil {
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

func (s *PresenceStatusController) GetPresenceStatusByID(ctx echo.Context) error {
	nip := ctx.Param("id")

	PresenceStatus, err := s.PresenceStatuservice.GetPresenceStatusByID(nip)
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
		Data:    PresenceStatus,
	})
}

func (s *PresenceStatusController) GetPresenceStatus(ctx echo.Context) error {
	PresenceStatus, err := s.PresenceStatuservice.GetPresenceStatus()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    PresenceStatus,
	})
}

func (s *PresenceStatusController) UpdatePresenceStatus(ctx echo.Context) error {
	payload := new(model.PresenceStatus)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.PresenceStatuservice.UpdatePresenceStatus(payload); err != nil {
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
