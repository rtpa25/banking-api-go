package service

import (
	"github.com/rtpa25/banking/domain"
	"github.com/rtpa25/banking/dto"
	"github.com/rtpa25/banking/errs"
)

//this is the interface that exposes functions used by the rest handlers
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status != "active" && status != "inactive" {
		status = ""
	}
	resp, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var finalResponse []dto.CustomerResponse

	for _, customerFromDB := range resp {
		customerDTOtype := customerFromDB.ToDto()
		finalResponse = append(finalResponse, customerDTOtype)
	}

	return finalResponse, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repo: repository,
	}
}
