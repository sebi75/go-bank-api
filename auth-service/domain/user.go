package domain

import (
	"banking-auth/dto"
	errs "banking-auth/error"
)

type User struct {
	Id  string
	Username string
	Password string
	CustomerId int
}

type UserRepository interface {
	FindById(string) (*User, *errs.AppError)
	Create(User) (*User, *errs.AppError)
}

func (u User) ToNewUserResponseDto() *dto.NewUserResponse {
	return &dto.NewUserResponse{
		UserId: u.Id,
		Username: u.Username,
		CustomerId: u.CustomerId,
	}
}