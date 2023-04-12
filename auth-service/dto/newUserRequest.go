package dto

type NewUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	CustomerId int `json:"customer_id"`
}