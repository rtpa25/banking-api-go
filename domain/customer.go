package domain

import "github.com/rtpa25/banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

//this is the interface that talks with the DB layer
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	//customer is pointer because it is optional
	ById(string) (*Customer, *errs.AppError)
}
