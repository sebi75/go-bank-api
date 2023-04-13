package service

import (
	"banking-auth/domain"
	"banking-auth/dto"
	errs "banking-auth/error"
)

type UserService interface {
	CreateUser(dto.RegisterRequest) (*dto.RegisterResponse, *errs.AppError)
}

type DefaultUserService struct {
	repo domain.AuthRepository
}

func (us DefaultUserService) CreateUser(req dto.RegisterRequest) (*dto.RegisterResponse, *errs.AppError) {
	req.Validate()

	//transform the NewUserRequest to User domain to save in db
	domainUser := domain.User{
		Id:         "",
		Username:   req.Username,
		Password:   req.Password,
		CustomerId: req.CustomerId,
	}

	return &dto.RegisterResponse{
		Token: "mockToken",
	}, nil
}

func NewUserService(repo domain.AuthRepository) DefaultUserService {
	return DefaultUserService{repo: repo}
}
