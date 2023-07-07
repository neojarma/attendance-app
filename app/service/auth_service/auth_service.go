package service

import (
	"errors"
	"presensi/helper"
	"presensi/model"
	repo "presensi/repository/auth_repository"
)

type AuthService struct {
	repo *repo.AuthRepository
}

func NewAuthService(r *repo.AuthRepository) *AuthService {
	return &AuthService{
		repo: r,
	}
}

func (s *AuthService) Auth(payload *model.LoginPayload) (*model.EmployeesJoinRole, error) {
	if payload.Username == "" || payload.Password == "" {
		return nil, errors.New("invalid body request")
	}

	employee, err := s.repo.Auth(payload.Username)
	if err != nil {
		return nil, err
	}

	if ok := helper.CheckPasswordHash(payload.Password, employee.Password); !ok {
		return nil, errors.New("invalid nip or password")
	}

	return employee, nil
}
