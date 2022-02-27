package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          1001,
			Name:        "Ronit",
			City:        "New Delhi",
			Zipcode:     "751012",
			DateOfBirth: "09-07-2002",
			Status:      "1",
		},
		{
			Id:          1002,
			Name:        "Yash",
			City:        "New Delhi",
			Zipcode:     "751012",
			DateOfBirth: "09-07-2002",
			Status:      "1",
		},
		{
			Id:          1003,
			Name:        "Nikhilesh",
			City:        "New Delhi",
			Zipcode:     "751012",
			DateOfBirth: "09-07-2002",
			Status:      "1",
		},
	}
	return CustomerRepositoryStub{
		customers: customers,
	}
}
