package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/work_times_service"

	"github.com/labstack/echo/v4"
)

type WorkTimesController struct {
	workTimeService *service.WorkTimesService
}

func NewWorkTimesController(ws *service.WorkTimesService) *WorkTimesController {
	return &WorkTimesController{
		workTimeService: ws,
	}
}

func (s *WorkTimesController) NewWorkTimes(ctx echo.Context) error {
	payload := new(model.WorkTimes)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.workTimeService.NewWorkTimes(payload); err != nil {
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

func (s *WorkTimesController) GetWorkTimesByID(ctx echo.Context) error {
	nip := ctx.Param("id")

	WorkTimes, err := s.workTimeService.GetWorkTimesByID(nip)
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
		Data:    WorkTimes,
	})
}

func (s *WorkTimesController) GetWorkTimes(ctx echo.Context) error {
	WorkTimess, err := s.workTimeService.GetWorkTimes()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    WorkTimess,
	})
}

func (s *WorkTimesController) UpdateWorkTimes(ctx echo.Context) error {
	payload := new(model.WorkTimes)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.workTimeService.UpdateWorkTimes(payload); err != nil {
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
