package domain

import "go-bank-api/errs"

type Account struct {
	AccountId string
	CustomerId string
	OpeningDate string
	AccountType string
	Amount float64
	Status string
}

type AccountRepository interface {
	GetAllAccount() ([]Account, *errs.AppError)
	Save(Account) (*Account, *errs.AppError)
}