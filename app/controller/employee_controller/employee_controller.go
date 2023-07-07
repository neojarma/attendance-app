package controller

import (
	"net/http"
	"presensi/model"
	service "presensi/service/employee_service"

	"github.com/labstack/echo/v4"
)

type EmployeeController struct {
	employeeService *service.EmployeeService
}

func NewEmployeeController(es *service.EmployeeService) *EmployeeController {
	return &EmployeeController{
		employeeService: es,
	}
}

func (s *EmployeeController) NewEmployee(ctx echo.Context) error {
	payload := new(model.Employees)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.employeeService.NewEmployee(payload); err != nil {
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

func (s *EmployeeController) GetEmployeeByNIP(ctx echo.Context) error {
	nip := ctx.Param("id")

	employee, err := s.employeeService.GetEmployeeByNIP(nip)
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
		Data:    employee,
	})
}

func (s *EmployeeController) GetEmployees(ctx echo.Context) error {
	employees, err := s.employeeService.GetEmployee()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, model.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, model.Response{
		Status:  "success",
		Message: "success get record",
		Data:    employees,
	})
}

func (s *EmployeeController) UpdateEmployee(ctx echo.Context) error {
	payload := new(model.Employees)

	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, model.Response{
			Status:  "failed",
			Message: "invalid body request",
		})
	}

	if err := s.employeeService.UpdateEmployee(payload); err != nil {
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

func (s *EmployeeController) DeleteEmployee(ctx echo.Context) error {
	nip := ctx.Param("id")

	if err := s.employeeService.DeleteEmployee(nip); err != nil {
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
		Message: "success delete record",
	})
}
