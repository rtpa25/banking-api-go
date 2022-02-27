package domain

import (
	"github.com/rtpa25/banking/dto"
	"github.com/rtpa25/banking/errs"
)

type Customer struct {
	Id          int64
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string
}

func (c *Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c *Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:      c.Id,
		Name:    c.Name,
		City:    c.City,
		Zipcode: c.Zipcode,
		Status:  c.statusAsText(),
	}
}

//this is the interface that talks with the DB layer
type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	//customer is pointer because it is optional
	ById(string) (*Customer, *errs.AppError)
}
