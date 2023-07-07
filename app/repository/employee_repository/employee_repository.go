package repo

import (
	"errors"
	"presensi/helper"
	"presensi/model"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	dB *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		dB: db,
	}
}

func (r *EmployeeRepository) NewEmployee(payload *model.Employees) error {
	err := r.dB.Create(payload).Error
	if err != nil {
		return helper.SQLErrorParser(err.Error())
	}

	return nil
}

func (r *EmployeeRepository) GetEmployeeByNIP(payload string) (*model.Employees, error) {
	model := new(model.Employees)
	err := r.dB.Omit("password").Where("nip = ?", payload).First(model).Error
	if err != nil {
		return nil, helper.SQLErrorParser(err.Error())
	}

	return model, err
}

func (r *EmployeeRepository) GetEmployees() ([]*model.Employees, error) {
	model := make([]*model.Employees, 0)
	err := r.dB.Omit("password").Find(&model).Error
	if err != nil {
		return nil, helper.SQLErrorParser(err.Error())
	}

	return model, err
}

func (r *EmployeeRepository) UpdateEmployee(payload *model.Employees) error {
	res := r.dB.Model(payload).Where("nip = ?", payload.Nip).Updates(payload)

	if res.Error != nil {
		return helper.SQLErrorParser(res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}

func (r *EmployeeRepository) DeleteEmployee(payload string) error {
	model := new(model.Employees)
	res := r.dB.Model(model).Where("nip = ?", payload).Update("is_active", false)

	if res.Error != nil {
		return helper.SQLErrorParser(res.Error.Error())
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}

func (r *EmployeeRepository) IsUsernameExist(username string) (bool, error) {
	model := new(model.Employees)
	res := r.dB.Model(model).Where("username = ?", username).Find(model)

	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected != 0, nil
}
