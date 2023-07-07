package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/absent_service"

	"github.com/labstack/echo/v4"
)

type AbsentController struct {
	AbsentService *service.AbsentService
}

func NewAbsentController(ps *service.AbsentService) *AbsentController {
	return &AbsentController{
		AbsentService: ps,
	}
}

func (c *AbsentController) InputAbsent(ctx echo.Context) error {
	payload := new(model.AbsentPayload)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := c.AbsentService.InputAbsent(payload); err != nil {
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

func (c *AbsentController) GetAbsentDaily(ctx echo.Context) error {

	date := ctx.QueryParam("date")
	from := ctx.QueryParam("from")
	to := ctx.QueryParam("to")

	res := make([]*model.Absent, 0)
	var err error

	if from != "" {
		res, err = c.AbsentService.GetAbsentDailyDateRange(from, to)
	} else {
		res, err = c.AbsentService.GetAbsentDaily(date)
	}

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    res,
	})
}
