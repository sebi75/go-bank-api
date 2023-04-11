package dto

type CustomerResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `json:"dateOfBirth"`
	Status      string `json:"status"`
}