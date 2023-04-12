package service

import (
	"banking-auth/domain"
	"banking-auth/dto"
	errs "banking-auth/error"
)

type UserService interface {
	CreateUser(dto.NewUserRequest) (*dto.NewUserResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.UserRepository
}

func (ud DefaultUserService) CreateUser(dto.NewUserRequest) (*dto.NewUserResponse, *errs.AppError) {
	return nil, nil // implement
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repo}
}
