package domain

import (
	errs "banking-auth/error"
)

type User struct {
	Id  string
	Username string
	Password string
	CustomerId int
	Role string
}

type AuthRepository interface {
	FindById(string) (*User, *errs.AppError)
	Create(User) (*User, *errs.AppError)
}
