package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/position_service"

	"github.com/labstack/echo/v4"
)

type PositionController struct {
	positionervice *service.Positionervice
}

func NewPositionController(es *service.Positionervice) *PositionController {
	return &PositionController{
		positionervice: es,
	}
}

func (s *PositionController) NewPosition(ctx echo.Context) error {
	payload := new(model.Position)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.positionervice.NewPosition(payload); err != nil {
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

func (s *PositionController) GetPositionByID(ctx echo.Context) error {
	nip := ctx.Param("id")

	Position, err := s.positionervice.GetPositionByID(nip)
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
		Data:    Position,
	})
}

func (s *PositionController) GetPosition(ctx echo.Context) error {
	Position, err := s.positionervice.GetPosition()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    Position,
	})
}

func (s *PositionController) UpdatePosition(ctx echo.Context) error {
	payload := new(model.Position)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.positionervice.UpdatePosition(payload); err != nil {
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
