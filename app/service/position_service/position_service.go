package service

import (
	"errors"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/position_repository"
)

type Positionervice struct {
	positionRepository *repo.PositionRepository
}

func NewPositionervice(rr *repo.PositionRepository) *Positionervice {
	return &Positionervice{
		positionRepository: rr,
	}
}

func (s *Positionervice) NewPosition(payload *model.Position) error {
	payload.PositionID = helper.GenerateRandString(5)
	return s.positionRepository.NewPosition(payload)
}

func (s *Positionervice) GetPositionByID(payload string) (*model.Position, error) {
	Position, err := s.positionRepository.GetPositionByID(payload)
	if err != nil {
		return nil, err
	}

	if Position.PositionID == "" {
		return nil, errors.New("there is no record with that id")
	}

	return Position, nil
}

func (s *Positionervice) GetPosition() ([]*model.Position, error) {
	return s.positionRepository.GetPosition()
}

func (s *Positionervice) UpdatePosition(payload *model.Position) error {
	return s.positionRepository.UpdatePosition(payload)
}
