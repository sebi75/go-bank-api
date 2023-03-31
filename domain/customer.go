package domain

import "go-bank-api/errs"

type Customer struct {
	Id string
	Name string
	City string
	Zipcode string
	DateofBirth string
	Status string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	FindById(customerId string) (*Customer, *errs.AppError)
	CreateCustomer(customer Customer) (*Customer, *errs.AppError)
}