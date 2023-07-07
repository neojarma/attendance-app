package repo

import (
	"errors"
	"presensi/model"

	"gorm.io/gorm"
)

type PositionRepository struct {
	dB *gorm.DB
}

func NewPositionRepository(db *gorm.DB) *PositionRepository {
	return &PositionRepository{
		dB: db,
	}
}

func (r *PositionRepository) NewPosition(payload *model.Position) error {
	return r.dB.Create(payload).Error
}

func (r *PositionRepository) GetPositionByID(payload string) (*model.Position, error) {
	model := new(model.Position)
	err := r.dB.Where("position_id = ?", payload).First(model).Error

	return model, err
}

func (r *PositionRepository) GetPosition() ([]*model.Position, error) {
	model := make([]*model.Position, 0)
	err := r.dB.Find(&model).Error

	return model, err
}

func (r *PositionRepository) UpdatePosition(payload *model.Position) error {
	res := r.dB.Model(payload).Where("position_id = ?", payload.PositionID).Updates(payload)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}
