package dto

type NewUserResponse struct {
	UserId string `json:"user_id"`
	Username string `json:"username"`
	CustomerId int `json:"customer_id"`
	Role string `json:"role"`
	Token string `json:"token"`
}