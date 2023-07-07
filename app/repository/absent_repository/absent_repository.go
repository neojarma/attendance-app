package absent_repository

import (
	"presensi/helper"
	"presensi/model"

	"gorm.io/gorm"
)

type AbsentRepository struct {
	dB *gorm.DB
}

func NewAbsentRepository(db *gorm.DB) *AbsentRepository {
	return &AbsentRepository{
		dB: db,
	}
}

func (r *AbsentRepository) InputAbsent(payload []*model.Absent) error {
	err := r.dB.Omit("name").Create(payload).Error
	if err != nil {
		return helper.SQLErrorParser(err.Error())
	}

	return nil
}

func (r *AbsentRepository) GetAbsentDaily(dateStr string) ([]*model.Absent, error) {
	model := make([]*model.Absent, 0)

	err := r.dB.
		Table("absents").
		Select("absents.absent_id, absents.nip, employees.name, absents.absent_date, presence_statuses.status_description as status_absent").
		Joins("JOIN presence_statuses ON absents.status_absent = presence_statuses.status_id").
		Joins("JOIN employees ON employees.nip = absents.nip").
		Where("absents.Absent_date = ? ", dateStr).
		Find(&model).
		Error

	return model, err
}

func (r *AbsentRepository) GetAbsentDailyDateRange(startDate, endDate string) ([]*model.Absent, error) {
	model := make([]*model.Absent, 0)

	err := r.dB.
		Table("absents").
		Select("absents.absent_id, absents.nip, employees.name, absents.absent_date, presence_statuses.status_description as status_absent").
		Joins("JOIN presence_statuses ON absents.status_absent = presence_statuses.status_id").
		Joins("JOIN employees ON employees.nip = absents.nip").
		Where("absents.absent_date BETWEEN ? AND ?", startDate, endDate).
		Find(&model).
		Error

	return model, err
}

func (r *AbsentRepository) GetAbsentMonthly(month, year string) ([]*model.AbsentAggregate, error) {
	model := make([]*model.AbsentAggregate, 0)

	err := r.dB.
		Table("absents").
		Select("absents.nip, absents.nip, employees.name, COUNT(CASE WHEN absents.status_absent = 'jMaJT' THEN 1 END) AS alpa, COUNT(CASE WHEN absents.status_absent = 'MjiqS' THEN 1 END) AS sakit, COUNT(CASE WHEN absents.status_absent = 'RQjtW' THEN 1 END) AS izin").
		Where("MONTH(absents.absent_date) = ? AND YEAR(absents.absent_date) = ?", month, year).
		Joins("JOIN employees ON employees.nip = absents.nip").
		Group("absents.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *AbsentRepository) GetAbsentYearly(year string) ([]*model.AbsentAggregate, error) {
	model := make([]*model.AbsentAggregate, 0)

	err := r.dB.
		Table("absents").
		Select("absents.nip, employees.name, COUNT(CASE WHEN absents.status_absent = 'jMaJT' THEN 1 END) AS alpa, COUNT(CASE WHEN absents.status_absent = 'MjiqS' THEN 1 END) AS sakit, COUNT(CASE WHEN absents.status_absent = 'RQjtW' THEN 1 END) AS izin").
		Where("YEAR(absents.absent_date) = ?", year).
		Joins("JOIN employees ON employees.nip = absents.nip").
		Group("absents.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *AbsentRepository) GetAbsentMonthlyWithNIP(month, year, nip string) ([]*model.AbsentAggregate, error) {
	model := make([]*model.AbsentAggregate, 0)

	err := r.dB.
		Table("absents").
		Select("absents.nip, employees.name, COUNT(CASE WHEN absents.status_absent = 'jMaJT' THEN 1 END) AS alpa, COUNT(CASE WHEN absents.status_absent = 'MjiqS' THEN 1 END) AS sakit, COUNT(CASE WHEN absents.status_absent = 'RQjtW' THEN 1 END) AS izin").
		Joins("JOIN employees ON employees.nip = absents.nip").
		Where("MONTH(absents.absent_date) = ? AND YEAR(absents.absent_date) = ? AND absents.nip = ?", month, year, nip).
		Group("absents.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *AbsentRepository) GetAbsentYearlyWithNIP(year, nip string) ([]*model.AbsentAggregate, error) {
	model := make([]*model.AbsentAggregate, 0)

	err := r.dB.
		Table("absents").
		Select("absents.nip, employees.name, COUNT(CASE WHEN absents.status_absent = 'jMaJT' THEN 1 END) AS alpa, COUNT(CASE WHEN absents.status_absent = 'MjiqS' THEN 1 END) AS sakit, COUNT(CASE WHEN absents.status_absent = 'RQjtW' THEN 1 END) AS izin").
		Joins("JOIN employees ON employees.nip = absents.nip").
		Where("YEAR(absents.absent_date) = ? AND absents.nip = ?", year, nip).
		Group("absents.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *AbsentRepository) AutoAbsent() error {
	return r.dB.Exec("INSERT INTO absents ( absent_id, nip, absent_date, status_absent ) SELECT SUBSTRING(CONVERT(VARCHAR(255), NEWID()), 1, 8), employees.nip, CONVERT(DATE, GETDATE()), 'jMaJT' FROM employees WHERE NOT EXISTS ( SELECT * FROM presences WHERE presences.nip = employees.nip ) AND employees.is_active = 1 AND NOT EXISTS ( SELECT * FROM holidays WHERE CONVERT ( DATE, GETDATE( ) ) = holidays.[date] ) AND datepart( WEEKDAY, GETDATE( ) ) NOT IN ( 1, 7 )").Error
}

func (r *AbsentRepository) GetUserAbsentWithNIP(nip string) ([]*model.Absent, error) {
	result := make([]*model.Absent, 0)

	return result, r.dB.Table("absents").Where("nip = ?", nip).Find(&result).Error
}
