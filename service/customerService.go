package service

import "go-bank-api/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (service DefaultCustomerService) GetAllCustomers() ([]domain.Customer, error) {
	return service.repo.FindAll()
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}