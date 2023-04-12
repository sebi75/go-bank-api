package service

import (
	"go-bank-api/domain"
	"go-bank-api/dto"
	"go-bank-api/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]domain.Customer, *errs.AppError)
	GetCustomerById(customerId string) (*dto.CustomerResponse, *errs.AppError)
	CreateCustomer(customer domain.Customer) (*domain.Customer, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers(status string) ([]domain.Customer, *errs.AppError) {
	return service.repo.FindAll(status)
}

func (service DefaultCustomerService) GetCustomerById(customerId string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := service.repo.FindById(customerId)
	if err != nil {
		return nil, err
	}
	return customer.ToDto(), nil
}

func (service DefaultCustomerService) CreateCustomer(customer domain.Customer) (*domain.Customer, *errs.AppError) {
	return service.repo.CreateCustomer(customer)
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
