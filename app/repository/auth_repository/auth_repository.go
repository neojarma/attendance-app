package repo

import (
	"errors"
	"presensi/model"

	"gorm.io/gorm"
)

type AuthRepository struct {
	dB *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *AuthRepository {
	return &AuthRepository{
		dB: db,
	}
}

func (r *AuthRepository) Auth(username string) (*model.EmployeesJoinRole, error) {
	model := &model.EmployeesJoinRole{
		Username: username,
	}

	res := r.dB.
		Table("employees").
		Select("employees.name, employees.nip, employees.position_id, roles.role_name AS role, employees.username, employees.password, employees.is_active").
		Joins("JOIN roles ON employees.role_id = roles.role_id").
		Where("username = ?", model.Username).
		Find(model)

	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("there is no record with that id")
	}

	return model, nil
}
