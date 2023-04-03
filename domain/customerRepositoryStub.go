package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (cr CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return cr.customers, nil
}

func (cr CustomerRepositoryStub) CreateCustomer(customer Customer) (*Customer, error) {
	cr.customers = append(cr.customers, customer)
	return &customer, nil
}

func (cr CustomerRepositoryStub) FindById(customerId string) (Customer, error) {
	var customer Customer
	for _, c := range cr.customers {
		if c.Id == customerId {
			customer = c
			break
		}
	}
	//check if customer is empty, and if yes return an error
	if (customer == Customer{}) {
		return Customer{}, nil
	}
	return customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Sebastian", "Berlin", "12345", "2000-01-01", "1"},
		{"1002", "John", "New York", "54321", "2000-01-01", "1"},
		{"1003", "Jane", "London", "11111", "2000-01-01", "1"},
	}
	return CustomerRepositoryStub{customers: customers}
}
