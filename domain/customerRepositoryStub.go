package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Rohan", City: "Karimganj", Zipcode: "788710", DateOfBirth: "1997-06-24", Status: "1"},
		{Id: "1002", Name: "Shalini", City: "Kolkata", Zipcode: "700039", DateOfBirth: "1997-05-13", Status: "1"},
	}

	return CustomerRepositoryStub{customers: customers}
}
