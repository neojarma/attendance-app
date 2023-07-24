package repo

import (
	"errors"
	"log"
	"presensi/helper"
	"presensi/model"

	"gorm.io/gorm"
)

type HolidayRepository struct {
	dB *gorm.DB
}

func NewHolidayRepository(db *gorm.DB) *HolidayRepository {
	return &HolidayRepository{
		dB: db,
	}
}

func (r *HolidayRepository) NewHoliday(payload *model.Holidays) error {
	err := r.dB.Create(payload).Error
	if err != nil {
		return helper.SQLErrorParser(err.Error())
	}

	return nil
}

func (r *HolidayRepository) GetHolidayByID(payload string) (*model.Holidays, error) {
	model := new(model.Holidays)
	err := r.dB.Where("holiday_id = ?", payload).First(model).Error

	return model, err
}

func (r *HolidayRepository) GetHoliday() ([]*model.Holidays, error) {
	model := make([]*model.Holidays, 0)
	err := r.dB.Find(&model).Error

	return model, err
}

func (r *HolidayRepository) UpdateHoliday(payload *model.Holidays) error {
	res := r.dB.Model(payload).Where("holiday_id = ?", payload.HolidayID).Updates(payload)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}

func (r *HolidayRepository) DeleteHoliday(payload string) error {
	model := &model.Holidays{
		HolidayID: payload,
	}

	res := r.dB.Where("holiday_id = ?", payload).Delete(model)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}

func (r *HolidayRepository) SeedPublicHoliday(payloads []*model.Holidays) error {
	if r.dB.Create(payloads).RowsAffected == 0 {
		return errors.New("there is no record with that id")
	}

	return nil
}

func (r *HolidayRepository) IsHolidayExistCurrentYear(year int) bool {
	model := new(model.Holidays)
	res := r.dB.Select("holiday_id").Where("YEAR(holidays.date) = ?", year).Find(model)

	if res.Error != nil {
		log.Println(res.Error)
		return false
	}

	return res.RowsAffected != 0
}

func (r *HolidayRepository) IsTodayAHoliday(date string) (bool, error) {
	model := new(model.Holidays)
	res := r.dB.Select("holiday_id").Where("holidays.date = ?", date).Find(model)

	if res.Error != nil {
		log.Println(res.Error)
		return false, res.Error
	}

	return res.RowsAffected != 0, nil
}
