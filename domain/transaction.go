package domain

import "github.com/rtpa25/banking/errs"

type Transaction struct {
	Id              int64   `db:"id"`
	AccountId       int64   `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionType string  `db:"transaction_type"`
	TransactionTime string  `db:"transaction_time"`
}

type TransactionRepository interface {
	Add(transaction Transaction) (*Transaction, *Account, *errs.AppError)
}
