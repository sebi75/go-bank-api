package domain

import (
	"go-bank-api/dto"
	"go-bank-api/errs"
)

type Account struct {
	AccountId string
	CustomerId string
	OpeningDate string
	AccountType string
	Amount float64
	Status string
}

func (a Account) ToNewAccountResponseDto() *dto.NewAccountResponse {
	return &dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}