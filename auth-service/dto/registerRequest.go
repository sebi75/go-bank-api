package dto

import errs "banking-auth/error"

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	CustomerId int `json:"customer_id"`
}

//Validate the basic fields of the request
//This can be improven by using a validation library
func (ur RegisterRequest) Validate() *errs.AppError {
	if ur.Username == "" {
		return errs.NewValidationError("invalid username")
	}
	if ur.Password == "" {
		return errs.NewValidationError("invalid password")
	}
	if ur.CustomerId == 0 {
		return errs.NewValidationError("invalid customer id")
	}
	return nil
}