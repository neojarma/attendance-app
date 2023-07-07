package service

import (
	"errors"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/presence_repository"
	service "presensi/service/absent_service"
	holiday_service "presensi/service/holiday_service"
	"strconv"
	"time"
)

type PresenceService struct {
	presenceRepo   *repo.PresenceRepository
	absentService  *service.AbsentService
	holidayService *holiday_service.HolidayService
}

func NewPresenceRepository(pr *repo.PresenceRepository, as *service.AbsentService, hs *holiday_service.HolidayService) *PresenceService {
	return &PresenceService{
		presenceRepo:   pr,
		absentService:  as,
		holidayService: hs,
	}
}

func (s *PresenceService) InputPresence(payload *model.PresencePayload) error {

	now := time.Now().Weekday()
	if now == time.Sunday || now == time.Saturday {
		return errors.New("today is off")
	}

	if !helper.IsDateValid(payload.PresenceDate) {
		return errors.New("invalid date format")
	}

	holiday, err := s.holidayService.IsTodayAHoliday(payload.PresenceDate)
	if err != nil {
		return err
	}

	if holiday {
		return errors.New("today is off")
	}

	exist, err := s.presenceRepo.IsRecordAlreadyExist(payload.PresenceDate, payload.Nip)
	if err != nil {
		return err
	}

	if exist {
		return errors.New("record already exist")
	}

	dbModel := model.Presence{
		PresenceId:   helper.GenerateRandString(10),
		WorkTimeId:   payload.WorkTimeId,
		Nip:          payload.Nip,
		WaktuMasuk:   payload.WaktuMasuk,
		PresenceDate: payload.PresenceDate,
	}

	return s.presenceRepo.InputPresence(&dbModel)
}

func (s *PresenceService) GetPresenceDaily(dateStr string) ([]*model.Presence, error) {
	if !helper.IsDateValid(dateStr) {
		return nil, errors.New("invalid date format")
	}

	res, err := s.presenceRepo.GetPresenceDaily(dateStr)
	if err != nil {
		return nil, err
	}

	return s.ParsingDailyPresence(res)
}

func (s *PresenceService) GetPresenceDailyDateRange(startDate, endDate string) ([]*model.Presence, error) {
	if !helper.IsDateValid(startDate) || !helper.IsDateValid(endDate) {
		return nil, errors.New("invalid date format")
	}

	res, err := s.presenceRepo.GetPresenceDailyDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return s.ParsingDailyPresence(res)
}

func (s *PresenceService) GetPresenceDailyReport(dateStr string) (string, error) {
	if !helper.IsDateValid(dateStr) {
		return "", errors.New("invalid date format")
	}

	tempRes, err := s.presenceRepo.GetPresenceDaily(dateStr)
	if err != nil {
		return "", err
	}

	finalRes, err := s.ParsingDailyPresence(tempRes)
	if err != nil {
		return "", err
	}

	link, err := helper.GetReport(finalRes, helper.GetPresenceDailyReportName(dateStr))
	if err != nil {
		return "", err
	}

	return link, nil
}

func (s *PresenceService) GetPresenceDailyDateRangeReport(startDate, endDate string) (string, error) {
	if !helper.IsDateValid(startDate) || !helper.IsDateValid(endDate) {
		return "", errors.New("invalid date format")
	}

	tempRes, err := s.presenceRepo.GetPresenceDailyDateRange(startDate, endDate)
	if err != nil {
		return "", err
	}

	finalRes, err := s.ParsingDailyPresence(tempRes)
	if err != nil {
		return "", err
	}

	link, err := helper.GetReport(finalRes, helper.GetPresenceDailyRangeReportName(startDate, endDate))
	if err != nil {
		return "", err
	}

	return link, nil
}

func (s *PresenceService) GetPresenceMonthly(month, year string) ([]*model.ReportAggregate, error) {

	_, errMonth := strconv.Atoi(month)
	_, errYear := strconv.Atoi(year)

	if errMonth != nil || errYear != nil {
		return nil, errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceMonthly(month, year)
	if err != nil {
		return nil, err
	}

	resAbsent, err := s.absentService.GetAbsentMonthly(month, year)
	if err != nil {
		return nil, err
	}

	return s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
}

func (s *PresenceService) GetPresenceMonthlyReport(month, year string) (string, error) {

	_, errMonth := strconv.Atoi(month)
	_, errYear := strconv.Atoi(year)

	if errMonth != nil || errYear != nil {
		return "", errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceMonthly(month, year)
	if err != nil {
		return "", err
	}

	resAbsent, err := s.absentService.GetAbsentMonthly(month, year)
	if err != nil {
		return "", err
	}

	joinRes, err := s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
	if err != nil {
		return "", err
	}

	link, err := helper.GetReport(joinRes, helper.GetPresenceMonthlyReportName(month, year))
	if err != nil {
		return "", err
	}

	return link, nil
}

func (s *PresenceService) GetPresenceYearly(year string) ([]*model.ReportAggregate, error) {
	_, errYear := strconv.Atoi(year)

	if errYear != nil {
		return nil, errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceYearly(year)
	if err != nil {
		return nil, err
	}

	resAbsent, err := s.absentService.GetAbsentYearly(year)
	if err != nil {
		return nil, err
	}

	return s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
}

func (s *PresenceService) GetPresenceYearlyReport(year string) (string, error) {
	_, errYear := strconv.Atoi(year)

	if errYear != nil {
		return "", errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceYearly(year)
	if err != nil {
		return "", err
	}

	resAbsent, err := s.absentService.GetAbsentYearly(year)
	if err != nil {
		return "", err
	}

	joinRes, err := s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
	if err != nil {
		return "", err
	}

	link, err := helper.GetReport(joinRes, helper.GetPresenceYearlyReportName(year))
	if err != nil {
		return "", err
	}

	return link, nil
}

func (s *PresenceService) GetPresenceMonthlyWithNIP(month, year, nip string) ([]*model.ReportAggregate, error) {
	_, errMonth := strconv.Atoi(month)
	_, errYear := strconv.Atoi(year)

	if errMonth != nil || errYear != nil {
		return nil, errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceMonthlyWithNIP(month, year, nip)
	if err != nil {
		return nil, err
	}

	resAbsent, err := s.absentService.GetAbsentMonthlyWithNIP(month, year, nip)
	if err != nil {
		return nil, err
	}

	return s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
}

func (s *PresenceService) GetPresenceMonthlyWithNIPReport(month, year, nip string) (string, error) {
	_, errMonth := strconv.Atoi(month)
	_, errYear := strconv.Atoi(year)

	if errMonth != nil || errYear != nil {
		return "", errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceMonthlyWithNIP(month, year, nip)
	if err != nil {
		return "", err
	}

	resAbsent, err := s.absentService.GetAbsentMonthlyWithNIP(month, year, nip)
	if err != nil {
		return "", err
	}

	joinRes, err := s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
	if err != nil {
		return "", err
	}

	name := joinRes[0].Name
	link, err := helper.GetReport(joinRes, helper.GetPresenceMonthlyNIPReportName(name, month, year))
	if err != nil {
		return "", err
	}

	return link, nil
}

func (s *PresenceService) GetPresenceYearlyWithNIP(year, nip string) ([]*model.ReportAggregate, error) {
	_, errYear := strconv.Atoi(year)

	if errYear != nil {
		return nil, errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceYearlyWithNIP(year, nip)
	if err != nil {
		return nil, err
	}

	resAbsent, err := s.absentService.GetAbsentYearlyWithNIP(year, nip)
	if err != nil {
		return nil, err
	}

	return s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
}

func (s *PresenceService) GetPresenceYearlyWithNIPReport(year, nip string) (string, error) {
	_, errYear := strconv.Atoi(year)

	if errYear != nil {
		return "", errors.New("invalid date format")
	}

	resPresence, err := s.presenceRepo.GetPresenceYearlyWithNIP(year, nip)
	if err != nil {
		return "", err
	}

	resAbsent, err := s.absentService.GetAbsentYearlyWithNIP(year, nip)
	if err != nil {
		return "", err
	}

	joinRes, err := s.parsingResultFromAbsentAndPresence(resAbsent, resPresence), nil
	if err != nil {
		return "", err
	}

	name := joinRes[0].Name
	link, err := helper.GetReport(joinRes, helper.GetPresenceYearlyNIPReportName(name, year))
	if err != nil {
		return "", err
	}

	return link, nil
}

func (s *PresenceService) UpdateJamPulang(payload *model.PresencePayload) error {
	dBModel := model.Presence{
		Nip:          payload.Nip,
		PresenceDate: payload.PresenceDate,
		WaktuKeluar:  payload.WaktuKeluar,
	}

	return s.presenceRepo.UpdateJamPulang(&dBModel)
}

func (s *PresenceService) parsingResultFromAbsentAndPresence(resAbsent []*model.AbsentAggregate, resPresence []*model.PresenceAggregate) []*model.ReportAggregate {
	tempResMap := make(map[string]*model.ReportAggregate)
	for _, v := range resPresence {
		tempResMap[v.Nip] = &model.ReportAggregate{
			Nip:         v.Nip,
			OnTime:      v.OnTime,
			Late:        v.Late,
			PulangTepat: v.PulangTepat,
			PulangCepat: v.PulangCepat,
			Name:        v.Name,
		}
	}

	for _, v := range resAbsent {
		if tempR, exist := tempResMap[v.Nip]; exist {
			tempR.Alpa = v.Alpa
			tempR.Sakit = v.Sakit
			tempR.Izin = v.Izin
		} else {
			tempResMap[v.Nip] = &model.ReportAggregate{
				Nip:   v.Nip,
				Name:  v.Name,
				Sakit: v.Sakit,
				Izin:  v.Izin,
				Alpa:  v.Alpa,
			}
		}
	}

	finalRes := make([]*model.ReportAggregate, 0)
	for _, v := range tempResMap {
		finalRes = append(finalRes, v)
	}

	return finalRes
}

func (s *PresenceService) AutoUpdateJamPulang() error {
	return s.presenceRepo.AutoUpdateJamPulang()
}

func (s *PresenceService) ParsingDailyPresence(data []*model.Presence) ([]*model.Presence, error) {
	for _, v := range data {
		var err error

		v.PresenceDate = helper.ParseDateOnly(v.PresenceDate)

		v.WaktuMasuk, err = helper.ParseTime(v.WaktuMasuk)
		if err != nil {
			return nil, err
		}

		v.WaktuKeluar, err = helper.ParseTime(v.WaktuKeluar)
		if err != nil {
			return nil, err
		}

	}

	return data, nil
}
