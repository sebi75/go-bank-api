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
	//call the repository method to save the user to the database
	user, err := us.repo.CreateUser(domainUser)
	if err != nil {
		return nil, err
	}

	token, tokenErr := user.GenerateToken()
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.RegisterResponse{
		Token: token,
	}, nil
}

func NewUserService(repo domain.AuthRepository) DefaultUserService {
	return DefaultUserService{repo: repo}
}
