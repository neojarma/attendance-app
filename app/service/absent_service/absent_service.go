package service

import (
	"errors"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/absent_repository"
	"strconv"
)

type AbsentService struct {
	AbsentRepo *repo.AbsentRepository
}

func NewAbsentService(pr *repo.AbsentRepository) *AbsentService {
	return &AbsentService{
		AbsentRepo: pr,
	}
}

func (s *AbsentService) InputAbsent(payload *model.AbsentPayload) error {
	dbPayload := make([]*model.Absent, 0)

	if len(payload.DateRange) != 0 {
		if len(payload.DateRange) != 2 {
			return errors.New("invalid date range format")
		}

		dateRange, err := helper.GenerateDateRange(payload.DateRange[0], payload.DateRange[1])
		if err != nil {
			return err
		}

		for _, v := range dateRange {
			if !helper.IsDateValid(v) {
				return errors.New("invalid date format")
			}

			res, err := s.AbsentRepo.GetUserAbsentWithNIP(payload.Nip)
			if err != nil {
				return err
			}

			if !s.IsDateExistInDB(res, v) {
				dbPayload = append(dbPayload, &model.Absent{
					AbsentId:     helper.GenerateRandString(10),
					Nip:          payload.Nip,
					AbsentDate:   v,
					StatusAbsent: payload.StatusAbsent,
					Keterangan:   payload.Keterangan,
				})
			}
		}
	} else {
		if !helper.IsDateValid(payload.Date) {
			return errors.New("invalid date format")
		}

		res, err := s.AbsentRepo.GetUserAbsentWithNIP(payload.Nip)
		if err != nil {
			return err
		}

		for _, v := range res {
			if helper.ParseDateOnly(v.AbsentDate) == payload.Date {
				return errors.New("user already has a record with that inputted date")
			}
		}

		dbPayload = append(dbPayload, &model.Absent{
			AbsentId:     helper.GenerateRandString(10),
			Nip:          payload.Nip,
			AbsentDate:   payload.Date,
			StatusAbsent: payload.StatusAbsent,
			Keterangan:   payload.Keterangan,
		})
	}

	if len(dbPayload) == 0 {
		return errors.New("all inputted date already exists")
	}

	return s.AbsentRepo.InputAbsent(dbPayload)
}

func (s *AbsentService) IsDateExistInDB(dbRes []*model.Absent, request string) bool {
	for _, v := range dbRes {
		if helper.ParseDateOnly(v.AbsentDate) == request {
			return true
		}
	}

	return false
}

func (s *AbsentService) GetAbsentDaily(dateStr string) ([]*model.Absent, error) {
	if !helper.IsDateValid(dateStr) {
		return nil, errors.New("invalid date format")
	}

	res, err := s.AbsentRepo.GetAbsentDaily(dateStr)
	if err != nil {
		return nil, err
	}

	return s.ParsingDate(res), nil
}

func (s *AbsentService) GetAbsentDailyDateRange(startDate, endDate string) ([]*model.Absent, error) {
	if !helper.IsDateValid(startDate) || !helper.IsDateValid(endDate) {
		return nil, errors.New("invalid date format")
	}

	res, err := s.AbsentRepo.GetAbsentDailyDateRange(startDate, endDate)
	if err != nil {
		return nil, err
	}

	return s.ParsingDate(res), nil
}

func (s *AbsentService) GetAbsentMonthly(month, year string) ([]*model.AbsentAggregate, error) {

	_, errMonth := strconv.Atoi(month)
	_, errYear := strconv.Atoi(year)

	if errMonth != nil || errYear != nil {
		return nil, errors.New("invalid date format")
	}

	return s.AbsentRepo.GetAbsentMonthly(month, year)
}

func (s *AbsentService) GetAbsentYearly(year string) ([]*model.AbsentAggregate, error) {
	_, errYear := strconv.Atoi(year)

	if errYear != nil {
		return nil, errors.New("invalid date format")
	}

	return s.AbsentRepo.GetAbsentYearly(year)
}

func (s *AbsentService) GetAbsentMonthlyWithNIP(month, year, nip string) ([]*model.AbsentAggregate, error) {
	_, errMonth := strconv.Atoi(month)
	_, errYear := strconv.Atoi(year)

	if errMonth != nil || errYear != nil {
		return nil, errors.New("invalid date format")
	}

	return s.AbsentRepo.GetAbsentMonthlyWithNIP(month, year, nip)
}

func (s *AbsentService) GetAbsentYearlyWithNIP(year, nip string) ([]*model.AbsentAggregate, error) {
	_, errYear := strconv.Atoi(year)

	if errYear != nil {
		return nil, errors.New("invalid date format")
	}

	return s.AbsentRepo.GetAbsentYearlyWithNIP(year, nip)
}

func (s *AbsentService) AutoAbsent() error {
	return s.AbsentRepo.AutoAbsent()
}

func (s *AbsentService) ParsingDate(res []*model.Absent) []*model.Absent {
	for _, v := range res {
		v.AbsentDate = helper.ParseDateOnly(v.AbsentDate)
	}

	return res
}
