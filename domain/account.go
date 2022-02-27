package domain

import (
	"github.com/rtpa25/banking/dto"
	"github.com/rtpa25/banking/errs"
)

type Account struct {
	Id          int64   `db:"id"`
	CustomerId  int64   `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      int     `db:"status"`
}

func (a *Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		Id: a.Id,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
