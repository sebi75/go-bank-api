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

func (us DefaultUserService) CreateUser(req dto.NewUserRequest) (*dto.NewUserResponse, *errs.AppError) {
	req.Validate()
	defaultTokenService := NewDefaultTokenService()

	//transform the NewUserRequest to User domain to save in db
	domainUser := domain.User{
		Id:         "",
		Username:   req.Username,
		Password:   req.Password,
		CustomerId: req.CustomerId,
	}

	//call the repository to save the user
	user, err := us.repo.Create(domainUser)
	if err != nil {
		return nil, err
	}

	userResponse := user.ToNewUserResponseDto()
	token, tokenErr  := defaultTokenService.GenerateToken(user)
	if tokenErr !=  nil {
		return nil, err
	}
	userResponse.Token = token

	return userResponse, nil
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo: repo}
}
