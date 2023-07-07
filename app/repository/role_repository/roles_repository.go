package repo

import (
	"errors"
	"presensi/model"

	"gorm.io/gorm"
)

type RolesRepository struct {
	dB *gorm.DB
}

func NewRolesRepository(db *gorm.DB) *RolesRepository {
	return &RolesRepository{
		dB: db,
	}
}

func (r *RolesRepository) NewRoles(payload *model.Roles) error {
	return r.dB.Create(payload).Error
}

func (r *RolesRepository) GetRolesByID(payload string) (*model.Roles, error) {
	model := new(model.Roles)
	err := r.dB.Where("role_id = ?", payload).First(model).Error

	return model, err
}

func (r *RolesRepository) GetRoles() ([]*model.Roles, error) {
	model := make([]*model.Roles, 0)
	err := r.dB.Find(&model).Error

	return model, err
}

func (r *RolesRepository) UpdateRoles(payload *model.Roles) error {
	res := r.dB.Model(payload).Where("role_id = ?", payload.RoleId).Updates(payload)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}
