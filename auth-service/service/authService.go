package service

import (
	"banking-auth/domain"
	"banking-auth/dto"
	errs "banking-auth/error"
)

type AuthService interface {
	CreateUser(dto.RegisterRequest) (*dto.RegisterResponse, *errs.AppError)
	LoginUser(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func (us DefaultAuthService) CreateUser(req dto.RegisterRequest) (*dto.RegisterResponse, *errs.AppError) {
	req.Validate()

	//transform the NewUserRequest to User domain to save in db
	domainUser := domain.User{
		Id:         "",
		Username:   req.Username,
		Password:   req.Password,
		CustomerId: req.CustomerId,
		Role: "USER",
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

func (us DefaultAuthService) LoginUser(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	req.Validate()

	//call the repository method to find the user by username
	user, err := us.repo.FindByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	//compare the password
	if user.Password != req.Password {
		return nil, errs.NewValidationError("invalid password")
	}

	token, tokenErr := user.GenerateToken()
	if tokenErr != nil {
		return nil, tokenErr
	}

	return &dto.LoginResponse{
		Token: token,
	}, nil
}

func NewUserService(repo domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{repo: repo}
}
