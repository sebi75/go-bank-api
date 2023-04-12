package dto

import "go-bank-api/errs"

type NewTransactionRequest struct {
	AccountId string `json:"account_id"`
	Amount float64 `json:"amount"`
	TransactionType string `json:"transaction_type"`
}

func (req NewTransactionRequest) Validate() *errs.AppError {
	if req.Amount < 0 {
		return errs.NewValidationError("Amount should be greater than 0")
	}

	if req.TransactionType != "DEPOSIT" && req.TransactionType != "WITHDRAWAL" {
		return errs.NewValidationError("Transaction type should be DEBIT or CREDIT")
	}

	return nil
}