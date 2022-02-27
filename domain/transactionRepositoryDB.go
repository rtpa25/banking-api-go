package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/rtpa25/banking/errs"
	"github.com/rtpa25/banking/logger"
)

type TransactionRepositoryDB struct {
	dbClient *sqlx.DB
}

func (d *TransactionRepositoryDB) Add(t Transaction) (*Transaction, *Account, *errs.AppError) {
	sqlToFetchExistingAccountBallance := `SELECT amount FROM accounts WHERE id=$1`
	var oldAccountBallance float64
	err := d.dbClient.Get(&oldAccountBallance, sqlToFetchExistingAccountBallance, t.AccountId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while fetching account ballance" + err.Error())
			return nil, nil, errs.NewNotFoundError("account not found")
		} else {
			logger.Error("Error while fetching account ballance" + err.Error())
			return nil, nil, errs.NewUnexpectedError("unexpected database error")
		}
	}

	if t.TransactionType == "withdrawl" {
		if t.Amount > oldAccountBallance {
			logger.Error("Insufficient ballance for withdrawl")
			return nil, nil, errs.NewValidationError("Insufficient ballance for withdrawl")
		} else {
			newTransaction, updatedAccount, err := d.processTransaction(t, oldAccountBallance)
			if err != nil {
				return nil, nil, errs.NewUnexpectedError("Unexpected error from database")
			}
			return newTransaction, updatedAccount, nil
		}
	} else if t.TransactionType == "deposit" {
		newTransaction, updatedAccount, err := d.processTransaction(t, oldAccountBallance)
		if err != nil {
			return nil, nil, errs.NewUnexpectedError("Unexpected error from database")
		}
		return newTransaction, updatedAccount, nil
	} else {
		logger.Error("Transaction Type can only be deposit or withdrawl")
		return nil, nil, errs.NewValidationError("Transaction Type can only be deposit or withdrawl")
	}
}

func (d *TransactionRepositoryDB) processTransaction(t Transaction, oldAccountBallance float64) (*Transaction, *Account, *errs.AppError) {
	sqlToInsertNewTransaction := `INSERT INTO transactions(account_id, transaction_type, amount) VALUES($1, $2, $3) RETURNING *`
	var transaction_id int64
	var account_id int64
	var transaction_time string
	var transaction_type string
	var transaction_amount float64
	result := d.dbClient.QueryRow(sqlToInsertNewTransaction, t.AccountId, t.TransactionType, t.Amount)
	err := result.Scan(&transaction_id, &account_id, &transaction_time, &transaction_type, &transaction_amount)
	if err != nil {
		logger.Error("Error while scanning for returned transaction: " + err.Error())
		return nil, nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	newTransaction := Transaction{
		Id:              transaction_id,
		AccountId:       account_id,
		TransactionType: transaction_type,
		TransactionTime: transaction_time,
		Amount:          transaction_amount,
	}

	var newBallance float64

	if t.TransactionType == "withdrawl" {
		newBallance = oldAccountBallance - t.Amount
	} else {
		newBallance = oldAccountBallance + t.Amount
	}

	sqlToUpdateAccount := `UPDATE accounts SET amount=$1 WHERE id=$2 RETURNING *`
	var id int64
	var customer_id int64
	var opening_date string
	var account_type string
	var amount float64
	var status int
	result = d.dbClient.QueryRow(sqlToUpdateAccount, newBallance, t.AccountId)
	err = result.Scan(&id, &customer_id, &opening_date, &account_type, &amount, &status)
	if err != nil {
		logger.Error("Error while scanning for returned account: " + err.Error())
		return nil, nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	updatedBankAccount := Account{
		Id:          id,
		CustomerId:  customer_id,
		OpeningDate: opening_date,
		AccountType: account_type,
		Amount:      amount,
		Status:      status,
	}
	return &newTransaction, &updatedBankAccount, nil
}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{dbClient: dbClient}
}
