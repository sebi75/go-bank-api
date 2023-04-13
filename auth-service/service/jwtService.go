package service

import (
	"banking-auth/domain"
	"banking-auth/logger"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenService interface {
	GenerateToken(user *domain.User) (string, error)
	VerifyToken(tokenString string) (*jwt.Token, error)
}

type DefaultTokenService struct {}

func (sv DefaultTokenService) GenerateToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{}
	claims["customer_id"] = user.CustomerId
	claims["user_id"] = user.Id
	claims["role"] = user.Role

	secretBytes := []byte(os.Getenv("JWT_SECRET"))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretBytes)

	if err != nil {
		logger.Error("Error while signing the token: " + err.Error())
		return "", err
	}
	logger.Info("tokenString: " + tokenString)
	return tokenString, nil
}

func (sv DefaultTokenService) VerifyToken(tokenString string) (*jwt.Token, error) {
	secretBytes := []byte(os.Getenv("JWT_SECRET"))
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretBytes, nil
	})

	if err != nil {
		logger.Error("Error while parsing the token: " + err.Error())
		return nil, err
	}

	return token, nil
}