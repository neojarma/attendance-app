package repo

import (
	"errors"
	"presensi/model"

	"gorm.io/gorm"
)

type WorkTimesRepository struct {
	dB *gorm.DB
}

func NewWorkTimeRepository(db *gorm.DB) *WorkTimesRepository {
	return &WorkTimesRepository{
		dB: db,
	}
}

func (r *WorkTimesRepository) NewWorkTime(payload *model.WorkTimes) error {
	return r.dB.Create(payload).Error
}

func (r *WorkTimesRepository) GetWorkTimesByID(payload string) (*model.WorkTimes, error) {
	model := new(model.WorkTimes)
	err := r.dB.Where("work_time_id = ?", payload).First(model).Error

	return model, err
}

func (r *WorkTimesRepository) GetWorkTimes() ([]*model.WorkTimes, error) {
	model := make([]*model.WorkTimes, 0)
	err := r.dB.Find(&model).Error

	return model, err
}

func (r *WorkTimesRepository) UpdateWorkTimes(payload *model.WorkTimes) error {
	res := r.dB.Model(payload).Where("work_time_id = ?", payload.WorkTimeId).Updates(payload)

	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}
