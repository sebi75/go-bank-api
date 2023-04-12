package dto

import "go-bank-api/errs"

type NewAccountRequest struct {
	CustomerId string `json:"customer_id"`
	AccountType string `json:"account_type"`
	Amount float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("Amount should be greater than 5000")
	}

	if r.AccountType != "SAVING" && r.AccountType != "CHECKING" {
		return errs.NewValidationError("Account type should be SAVING or CHECKING")
	}

	return nil
}