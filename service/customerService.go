package service

import (
	"go-bank-api/domain"
	"go-bank-api/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
	GetCustomerById(customerId string) (*domain.Customer, *errs.AppError)
	CreateCustomer(customer domain.Customer) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return service.repo.FindAll()
}

func (service DefaultCustomerService) GetCustomerById(customerId string) (*domain.Customer, *errs.AppError) {
	return service.repo.FindById(customerId)
}

func (service DefaultCustomerService) CreateCustomer(customer domain.Customer) (*domain.Customer, *errs.AppError) {
	return service.repo.CreateCustomer(customer)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}