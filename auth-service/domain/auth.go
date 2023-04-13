package domain

import (
	errs "banking-auth/error"
	"os"

	"github.com/golang-jwt/jwt/v5"
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
	CreateUser(User) (*User, *errs.AppError)
}

func (u User) GenerateToken() (string, *errs.AppError) {
	claims := jwt.MapClaims{}
	claims["username"] = u.Username
	claims["role"] = u.Role
	claims["customer_id"] = u.CustomerId

	secret := []byte(os.Getenv("JWT_SECRET"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", errs.NewUnexpectedError("Unexpected error from JWT library")
	}
	return tokenString, nil
}
