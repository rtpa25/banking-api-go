package service

import (
	"time"

	"github.com/rtpa25/banking/domain"
	"github.com/rtpa25/banking/dto"
	"github.com/rtpa25/banking/errs"
)

type TransactionService interface {
	AddNewTransaction(dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repo domain.TransactionRepository
}

func (d DefaultTransactionService) AddNewTransaction(req dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	t := domain.Transaction{
		Id:              0,
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionTime: time.Now().Format("2006-01-02T15:04:05Z07:00"),
	}
	newTransaction, updatedAccount, err := d.repo.Add(t)
	if err != nil {
		return nil, err
	}

	return &dto.NewTransactionResponse{Id: newTransaction.Id, UpdatedBankBallance: updatedAccount.Amount}, nil
}

func NewTransactionService(repo domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{
		repo: repo,
	}
}
