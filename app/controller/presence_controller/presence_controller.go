package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/presence_service"

	"github.com/labstack/echo/v4"
)

type PresenceController struct {
	presenceService *service.PresenceService
}

func NewPresenceRepository(ps *service.PresenceService) *PresenceController {
	return &PresenceController{
		presenceService: ps,
	}
}

func (c *PresenceController) InputPresence(ctx echo.Context) error {
	payload := new(model.PresencePayload)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := c.presenceService.InputPresence(payload); err != nil {
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

func (c *PresenceController) GetPresenceDaily(ctx echo.Context) error {

	date := ctx.QueryParam("date")
	from := ctx.QueryParam("from")
	to := ctx.QueryParam("to")

	res := make([]*model.Presence, 0)
	var err error

	if from != "" {
		res, err = c.presenceService.GetPresenceDailyDateRange(from, to)
	} else {
		res, err = c.presenceService.GetPresenceDaily(date)
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

func (c *PresenceController) GetPresenceDailyReport(ctx echo.Context) error {

	date := ctx.QueryParam("date")
	from := ctx.QueryParam("from")
	to := ctx.QueryParam("to")

	res := ""
	var err error

	if from != "" {
		res, err = c.presenceService.GetPresenceDailyDateRangeReport(from, to)
	} else {
		res, err = c.presenceService.GetPresenceDailyReport(date)
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

func (c *PresenceController) GetPresenceMonthly(ctx echo.Context) error {
	month := ctx.QueryParam("month")
	year := ctx.QueryParam("year")
	res, err := c.presenceService.GetPresenceMonthly(month, year)
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

func (c *PresenceController) GetPresenceMonthlyReport(ctx echo.Context) error {
	month := ctx.QueryParam("month")
	year := ctx.QueryParam("year")
	res, err := c.presenceService.GetPresenceMonthlyReport(month, year)
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

func (c *PresenceController) GetPresenceYearly(ctx echo.Context) error {
	year := ctx.QueryParam("year")
	res, err := c.presenceService.GetPresenceYearly(year)
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

func (c *PresenceController) GetPresenceYearlyReport(ctx echo.Context) error {
	year := ctx.QueryParam("year")
	res, err := c.presenceService.GetPresenceYearlyReport(year)
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

func (c *PresenceController) GetPresenceMonthlyWithNIP(ctx echo.Context) error {
	month := ctx.QueryParam("month")
	year := ctx.QueryParam("year")
	nip := ctx.QueryParam("nip")
	res, err := c.presenceService.GetPresenceMonthlyWithNIP(month, year, nip)
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

func (c *PresenceController) GetPresenceMonthlyWithNIPReport(ctx echo.Context) error {
	month := ctx.QueryParam("month")
	year := ctx.QueryParam("year")
	nip := ctx.QueryParam("nip")
	res, err := c.presenceService.GetPresenceMonthlyWithNIPReport(month, year, nip)
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

func (c *PresenceController) GetPresenceYearlyWithNIP(ctx echo.Context) error {
	year := ctx.QueryParam("year")
	nip := ctx.QueryParam("nip")
	res, err := c.presenceService.GetPresenceYearlyWithNIP(year, nip)
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

func (c *PresenceController) GetPresenceYearlyWithNIPReport(ctx echo.Context) error {
	year := ctx.QueryParam("year")
	nip := ctx.QueryParam("nip")
	res, err := c.presenceService.GetPresenceYearlyWithNIPReport(year, nip)
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

func (c *PresenceController) UpdateJamPulang(ctx echo.Context) error {
	payload := new(model.PresencePayload)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := c.presenceService.UpdateJamPulang(payload); err != nil {
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
