package domain

import errs "banking-auth/error"

type User struct {
	Id  int
	Username string
	Password string
	CustomerId int
}

type UserRepository interface {
	FindById(string) (*User, *errs.AppError)
	Create(User) (*User, *errs.AppError)
}