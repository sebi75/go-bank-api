package dto

import errs "banking-auth/error"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (lr LoginRequest) Validate() *errs.AppError {
	//basic validation
	if lr.Username == "" {
		return errs.NewValidationError("invalid username")
	}
	if lr.Password == "" {
		return errs.NewValidationError("invalid password")
	}
	return nil
}