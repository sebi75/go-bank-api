package domain

import (
	"go-bank-api/dto"
	"go-bank-api/errs"
)

type Transaction struct {
	Id string
	AccountId string
	Amount float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
	CreateTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToNewTransactionResponseDto() *dto.NewTransactionResponse {
	return &dto.NewTransactionResponse{TransactionId: t.Id}
}
