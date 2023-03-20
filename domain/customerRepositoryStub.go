package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Sebastian", "Berlin", "12345", "2000-01-01", "1"},
		{"1002", "John", "New York", "54321", "2000-01-01", "1"},
		{"1003", "Jane", "London", "11111", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}