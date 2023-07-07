package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/holiday_service"

	"github.com/labstack/echo/v4"
)

type HolidayController struct {
	holidayService *service.HolidayService
}

func NewHolidayController(hs *service.HolidayService) *HolidayController {
	return &HolidayController{
		holidayService: hs,
	}
}

func (s *HolidayController) NewHoliday(ctx echo.Context) error {
	payload := new(model.Holidays)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.holidayService.NewHoliday(payload); err != nil {
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

func (s *HolidayController) GetHolidayByID(ctx echo.Context) error {
	nip := ctx.Param("id")

	Holiday, err := s.holidayService.GetHolidayByID(nip)
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
		Data:    Holiday,
	})
}

func (s *HolidayController) GetHolidays(ctx echo.Context) error {
	Holiday, err := s.holidayService.GetHoliday()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    Holiday,
	})
}

func (s *HolidayController) UpdateHoliday(ctx echo.Context) error {
	payload := new(model.Holidays)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.holidayService.UpdateHoliday(payload); err != nil {
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

func (s *HolidayController) DeleteHoliday(ctx echo.Context) error {
	payload := ctx.Param("id")

	if err := s.holidayService.DeleteHoliday(payload); err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success delete record",
	})
}
