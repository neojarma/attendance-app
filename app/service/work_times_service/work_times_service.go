package service

import (
	"errors"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/work_times_repository"
)

type WorkTimesService struct {
	workTimesRepository *repo.WorkTimesRepository
}

func NewWorkTimesService(wr *repo.WorkTimesRepository) *WorkTimesService {
	return &WorkTimesService{
		workTimesRepository: wr,
	}
}

func (s *WorkTimesService) NewWorkTimes(payload *model.WorkTimes) error {
	payload.WorkTimeId = helper.GenerateRandString(5)
	return s.workTimesRepository.NewWorkTime(payload)
}

func (s *WorkTimesService) GetWorkTimesByID(payload string) (*model.WorkTimes, error) {
	WorkTimes, err := s.workTimesRepository.GetWorkTimesByID(payload)
	if err != nil {
		return nil, err
	}

	if WorkTimes.WorkTimeId == "" {
		return nil, errors.New("there is no record with that id")
	}

	parsedJamMasuk, err := helper.ParseTime(WorkTimes.JamMasuk)
	if err != nil {
		return nil, err
	}

	parsedJamKeluar, err := helper.ParseTime(WorkTimes.JamKeluar)
	if err != nil {
		return nil, err
	}

	WorkTimes.JamMasuk = parsedJamMasuk
	WorkTimes.JamKeluar = parsedJamKeluar
	return WorkTimes, nil
}

func (s *WorkTimesService) GetWorkTimes() ([]*model.WorkTimes, error) {
	res, err := s.workTimesRepository.GetWorkTimes()
	if err != nil {
		return nil, err
	}

	for _, v := range res {
		parsedJamMasuk, err := helper.ParseTime(v.JamMasuk)
		if err != nil {
			return nil, err
		}

		parsedJamKeluar, err := helper.ParseTime(v.JamKeluar)
		if err != nil {
			return nil, err
		}

		v.JamMasuk = parsedJamMasuk
		v.JamKeluar = parsedJamKeluar
	}

	return res, nil
}

func (s *WorkTimesService) UpdateWorkTimes(payload *model.WorkTimes) error {
	return s.workTimesRepository.UpdateWorkTimes(payload)
}
