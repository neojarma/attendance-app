package service

import (
	"errors"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/role_repository"
)

type Roleservice struct {
	rolesRepository *repo.RolesRepository
}

func NewRoleservice(rr *repo.RolesRepository) *Roleservice {
	return &Roleservice{
		rolesRepository: rr,
	}
}

func (s *Roleservice) NewRoles(payload *model.Roles) error {
	payload.RoleId = helper.GenerateRandString(5)
	return s.rolesRepository.NewRoles(payload)
}

func (s *Roleservice) GetRolesByID(payload string) (*model.Roles, error) {
	Roles, err := s.rolesRepository.GetRolesByID(payload)
	if err != nil {
		return nil, err
	}

	if Roles.RoleId == "" {
		return nil, errors.New("there is no record with that id")
	}

	return Roles, nil
}

func (s *Roleservice) GetRoles() ([]*model.Roles, error) {
	return s.rolesRepository.GetRoles()
}

func (s *Roleservice) UpdateRoles(payload *model.Roles) error {
	return s.rolesRepository.UpdateRoles(payload)
}
