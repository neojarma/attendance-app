package service

import (
	"errors"
	"log"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/employee_repository"
)

type EmployeeService struct {
	employeeRepository *repo.EmployeeRepository
}

func NewEmployeeService(er *repo.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		employeeRepository: er,
	}
}

func (s *EmployeeService) NewEmployee(payload *model.Employees) error {
	exist, err := s.employeeRepository.IsUsernameExist(payload.Username)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("username already exist")
	}

	hashedPass, err := helper.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	payload.Password = hashedPass
	return s.employeeRepository.NewEmployee(payload)
}

func (s *EmployeeService) GetEmployeeByNIP(payload string) (*model.Employees, error) {
	employee, err := s.employeeRepository.GetEmployeeByNIP(payload)
	if err != nil {
		return nil, err
	}

	if employee.Nip == "" {
		return nil, errors.New("there is no record with that id")
	}

	return employee, nil
}

func (s *EmployeeService) GetEmployee() ([]*model.Employees, error) {
	return s.employeeRepository.GetEmployees()
}

func (s *EmployeeService) UpdateEmployee(payload *model.Employees) error {
	if payload.Password != "" {
		hashedPass, err := helper.HashPassword(payload.Password)
		if err != nil {
			log.Println(err)
			return err
		}

		payload.Password = hashedPass
	}

	return s.employeeRepository.UpdateEmployee(payload)
}

func (s *EmployeeService) DeleteEmployee(payload string) error {
	return s.employeeRepository.DeleteEmployee(payload)
}
