package repo

import (
	"presensi/helper"
	"presensi/model"

	"gorm.io/gorm"
)

type PresenceRepository struct {
	dB *gorm.DB
}

func NewPresenceRepository(db *gorm.DB) *PresenceRepository {
	return &PresenceRepository{
		dB: db,
	}
}

func (r *PresenceRepository) InputPresence(payload *model.Presence) error {
	err := r.dB.Omit("status_masuk", "waktu_keluar", "status_keluar", "name").Create(payload).Error
	if err != nil {
		return helper.SQLErrorParser(err.Error())
	}

	return nil
}

func (r *PresenceRepository) GetPresenceDaily(dateStr string) ([]*model.Presence, error) {
	model := make([]*model.Presence, 0)

	err := r.dB.
		Select("presences.presence_id, presences.work_time_id, presences.nip, employees.name, presences.waktu_masuk, ps1.status_description AS status_masuk, presences.waktu_keluar, ps2.status_description AS status_keluar, presences.presence_date").
		Joins("JOIN presence_statuses ps1 ON presences.status_masuk = ps1.status_id").
		Joins("JOIN presence_statuses ps2 ON presences.status_keluar = ps2.status_id").
		Joins("JOIN employees ON employees.nip = presences.nip").
		Where("presences.presence_date = ? ", dateStr).
		Find(&model).
		Error

	return model, err
}

func (r *PresenceRepository) GetPresenceDailyDateRange(startDate, endDate string) ([]*model.Presence, error) {
	model := make([]*model.Presence, 0)

	err := r.dB.
		Select("presences.presence_id, presences.work_time_id, presences.nip, employees.name, presences.waktu_masuk, ps1.status_description AS status_masuk, presences.waktu_keluar, ps2.status_description AS status_keluar, presences.presence_date").
		Joins("JOIN presence_statuses ps1 ON presences.status_masuk = ps1.status_id").
		Joins("JOIN presence_statuses ps2 ON presences.status_keluar = ps2.status_id").
		Joins("JOIN employees ON employees.nip = presences.nip").
		Where("presences.presence_date BETWEEN ? AND ? ", startDate, endDate).
		Find(&model).
		Error

	return model, err
}

func (r *PresenceRepository) GetPresenceMonthly(month, year string) ([]*model.PresenceAggregate, error) {
	model := make([]*model.PresenceAggregate, 0)

	err := r.dB.
		Table("presences").
		Select("presences.nip, employees.name, COUNT(CASE WHEN presences.status_masuk = 'TppPt' THEN 1 END) AS on_time, COUNT(CASE WHEN presences.status_masuk = 'EQdio' THEN 1 END) AS late, COUNT(CASE WHEN presences.status_keluar = 'RiKuO' THEN 1 END) AS pulang_tepat, COUNT(CASE WHEN presences.status_keluar = 'raxbJ' THEN 1 END) AS pulang_cepat").
		Joins("JOIN employees ON employees.nip = presences.nip").
		Where("MONTH(presences.presence_date) = ? AND YEAR(presences.presence_date) = ?", month, year).
		Group("presences.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *PresenceRepository) GetPresenceYearly(year string) ([]*model.PresenceAggregate, error) {
	model := make([]*model.PresenceAggregate, 0)

	err := r.dB.
		Table("presences").
		Select("presences.nip, employees.name, COUNT(CASE WHEN presences.status_masuk = 'TppPt' THEN 1 END) AS on_time, COUNT(CASE WHEN presences.status_masuk = 'EQdio' THEN 1 END) AS late, COUNT(CASE WHEN presences.status_keluar = 'RiKuO' THEN 1 END) AS pulang_tepat, COUNT(CASE WHEN presences.status_keluar = 'raxbJ' THEN 1 END) AS pulang_cepat").
		Joins("JOIN employees ON employees.nip = presences.nip").
		Where("YEAR(presences.presence_date) = ?", year).
		Group("presences.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *PresenceRepository) GetPresenceMonthlyWithNIP(month, year, nip string) ([]*model.PresenceAggregate, error) {
	model := make([]*model.PresenceAggregate, 0)

	err := r.dB.
		Table("presences").
		Select("presences.nip, employees.name, COUNT(CASE WHEN presences.status_masuk = 'TppPt' THEN 1 END) AS on_time, COUNT(CASE WHEN presences.status_masuk = 'EQdio' THEN 1 END) AS late, COUNT(CASE WHEN presences.status_keluar = 'RiKuO' THEN 1 END) AS pulang_tepat, COUNT(CASE WHEN presences.status_keluar = 'raxbJ' THEN 1 END) AS pulang_cepat").
		Joins("JOIN employees ON employees.nip = presences.nip").
		Where("MONTH(presences.presence_date) = ? AND YEAR(presences.presence_date) = ? AND presences.nip = ?", month, year, nip).
		Group("presences.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *PresenceRepository) GetPresenceYearlyWithNIP(year, nip string) ([]*model.PresenceAggregate, error) {
	model := make([]*model.PresenceAggregate, 0)

	err := r.dB.
		Table("presences").
		Select("presences.nip, employees.name, COUNT(CASE WHEN presences.status_masuk = 'TppPt' THEN 1 END) AS on_time, COUNT(CASE WHEN presences.status_masuk = 'EQdio' THEN 1 END) AS late, COUNT(CASE WHEN presences.status_keluar = 'RiKuO' THEN 1 END) AS pulang_tepat, COUNT(CASE WHEN presences.status_keluar = 'raxbJ' THEN 1 END) AS pulang_cepat").
		Joins("JOIN employees ON employees.nip = presences.nip").
		Where("YEAR(presences.presence_date) = ? AND presences.nip = ?", year, nip).
		Group("presences.nip, employees.name").
		Scan(&model).
		Error

	return model, err
}

func (r *PresenceRepository) UpdateJamPulang(payload *model.Presence) error {
	return r.dB.Model(payload).Where("nip = ? AND presence_date = ?", payload.Nip, payload.PresenceDate).Updates(payload).Error
}

func (r *PresenceRepository) AutoUpdateJamPulang() error {
	return r.dB.Exec("UPDATE presences SET waktu_keluar = '15:30:00', status_keluar = 'RiKuO' WHERE presence_date = CONVERT ( DATE, GETDATE() )").Error
}

func (r *PresenceRepository) IsRecordAlreadyExist(date, nip string) (bool, error) {
	model := new(model.Presence)
	res := r.dB.Select("nip").Where("nip = ? and presence_date = ?", nip, date).Find(model)
	if res.Error != nil {
		return false, res.Error
	}

	return res.RowsAffected != 0, nil
}
