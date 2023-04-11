package service

import (
	"go-bank-api/domain"
	"go-bank-api/dto"
	"go-bank-api/errs"
	"time"
)

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (s DefaultTransactionService) NewTransaction(request dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}

	transaction := domain.Transaction{
		Id: "",
		AccountId: request.AccountId,
		Amount: request.Amount,
		TransactionType: request.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	newTransaction, appErr := s.repo.CreateTransaction(transaction)
	if appErr != nil {
		return nil, appErr
	}

	return newTransaction.ToNewTransactionResponseDto(), nil
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{repo: repo}
}
