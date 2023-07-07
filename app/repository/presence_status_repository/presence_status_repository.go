package repo

import (
	"errors"
	"presensi/model"

	"gorm.io/gorm"
)

type PresenceStatusRepository struct {
	dB *gorm.DB
}

func NewPresenceStatusRepository(db *gorm.DB) *PresenceStatusRepository {
	return &PresenceStatusRepository{
		dB: db,
	}
}

func (r *PresenceStatusRepository) NewPresenceStatus(payload *model.PresenceStatus) error {
	return r.dB.Create(payload).Error
}

func (r *PresenceStatusRepository) GetPresenceStatusByID(payload string) (*model.PresenceStatus, error) {
	model := new(model.PresenceStatus)
	err := r.dB.Where("status_id = ?", payload).First(model).Error

	return model, err
}

func (r *PresenceStatusRepository) GetPresenceStatus() ([]*model.PresenceStatus, error) {
	model := make([]*model.PresenceStatus, 0)
	err := r.dB.Find(&model).Error

	return model, err
}

func (r *PresenceStatusRepository) UpdatePresenceStatus(payload *model.PresenceStatus) error {
	res := r.dB.Model(payload).Where("status_id = ?", payload.StatusId).Updates(payload)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}
