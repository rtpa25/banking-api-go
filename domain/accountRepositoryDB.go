package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/rtpa25/banking/errs"
	"github.com/rtpa25/banking/logger"
)

type AccountRepositoryDB struct {
	dbClient *sqlx.DB
}

func (d *AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `INSERT INTO accounts(customer_id, account_type, amount) VALUES ($1, $2, $3) RETURNING id`
	res := d.dbClient.QueryRow(sqlInsert, a.CustomerId, a.AccountType, a.Amount)
	var id int64
	err := res.Scan(&id)
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.Id = id
	return &a, nil
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient: dbClient}
}
