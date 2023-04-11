package service

import (
	"go-bank-api/domain"
	"go-bank-api/dto"
	"go-bank-api/errs"
	"time"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (as DefaultAccountService) NewAccount(request dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := request.Validate()
	if err != nil {
		return nil, err
	}
	
	account := domain.Account{
		AccountId: "",
		CustomerId: request.CustomerId,
		AccountType: request.AccountType,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		Amount: request.Amount,
		Status: "1",
	}
	newAccount, appErr := as.repo.Save(account)
	if appErr != nil {
		return nil, appErr
	}
	return newAccount.ToNewAccountResponseDto(), nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo: repo}
}