package service

import (
	"banking-auth/domain"
	"banking-auth/logger"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *domain.User) (string, error) {
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