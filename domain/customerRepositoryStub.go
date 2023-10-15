package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {

	customers := []Customer{
		{"1001", "Sabari", "New Delhi", "110011", "2002-06-04", "1"},
		{"1002", "Deepi", "Chennai", "110011", "2002-06-04", "1"},
	}
	return CustomerRepositoryStub{customers}
}
