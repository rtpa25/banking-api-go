package domain

import "github.com/rtpa25/banking/errs"

//db tag is required when doing structscan with sqlx
type Customer struct {
	Id          string `db:"id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

//this is the interface that talks with the DB layer
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	//customer is pointer because it is optional
	ById(string) (*Customer, *errs.AppError)
}
