package service

import (
	"errors"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/presence_status_repository"
)

type PresenceStatuservice struct {
	presenceStatusRepository *repo.PresenceStatusRepository
}

func NewPresenceStatuservice(pr *repo.PresenceStatusRepository) *PresenceStatuservice {
	return &PresenceStatuservice{
		presenceStatusRepository: pr,
	}
}

func (s *PresenceStatuservice) NewPresenceStatus(payload *model.PresenceStatus) error {
	payload.StatusId = helper.GenerateRandString(5)
	return s.presenceStatusRepository.NewPresenceStatus(payload)
}

func (s *PresenceStatuservice) GetPresenceStatusByID(payload string) (*model.PresenceStatus, error) {
	PresenceStatus, err := s.presenceStatusRepository.GetPresenceStatusByID(payload)
	if err != nil {
		return nil, err
	}

	if PresenceStatus.StatusId == "" {
		return nil, errors.New("there is no record with that id")
	}

	return PresenceStatus, nil
}

func (s *PresenceStatuservice) GetPresenceStatus() ([]*model.PresenceStatus, error) {
	return s.presenceStatusRepository.GetPresenceStatus()
}

func (s *PresenceStatuservice) UpdatePresenceStatus(payload *model.PresenceStatus) error {
	return s.presenceStatusRepository.UpdatePresenceStatus(payload)
}
