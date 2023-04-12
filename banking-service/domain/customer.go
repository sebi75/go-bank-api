package domain

import (
	"go-bank-api/dto"
	"go-bank-api/errs"
)

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	Zipcode     string `json:"zipcode"`
	DateofBirth string `db:"date_of_birth"` // db tag is used by sqlx
	Status      string `json:"status"`
}

func (customer Customer) statusAsText() string {
	statusAsText := "active"

	if customer.Status == "0" {
		statusAsText = "inactive"
	}

	return statusAsText
}

func (customer *Customer) ToDto() *dto.CustomerResponse {
	return &dto.CustomerResponse{
		Id:          customer.Id,
		Name:        customer.Name,
		City:        customer.City,
		Zipcode:     customer.Zipcode,
		DateofBirth: customer.DateofBirth,
		Status:      customer.statusAsText(),
	}
}

type CustomerRepository interface {
	FindAll(string) ([]Customer, *errs.AppError)
	FindById(customerId string) (*Customer, *errs.AppError)
	CreateCustomer(customer Customer) (*Customer, *errs.AppError)
}
