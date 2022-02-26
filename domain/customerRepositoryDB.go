package domain

import (
	"database/sql"
	"strconv"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rtpa25/banking/errs"
	"github.com/rtpa25/banking/logger"
	"github.com/spf13/viper"
)

type CustomerRepositoryDB struct {
	dbClient *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSQL string
	var err error
	customers := make([]Customer, 0)

	if status == "" {
		findAllSQL = `SELECT id, name, city, zipcode, date_of_birth, status from customers`
		err = d.dbClient.Select(&customers, findAllSQL)
	} else {
		findAllSQL = `SELECT id, name, city, zipcode, date_of_birth, status from customers WHERE status = $1`
		if status == "active" {
			err = d.dbClient.Select(&customers, findAllSQL, 1)
		} else if status == "inactive" {
			err = d.dbClient.Select(&customers, findAllSQL, 0)
		}
	}

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while querying customer table " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}

	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findUserSQL := `SELECT id, name, city, zipcode, date_of_birth, status from customers WHERE id = $1`
	intId, err := strconv.Atoi(id)
	if err != nil {
		logger.Error(err.Error())
	}
	var c Customer
	err = d.dbClient.Get(&c, findUserSQL, intId)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
	}
	return &c, nil
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := viper.Get("DB_URL")
	db, err := sqlx.Open("postgres", connStr.(string))
	if err != nil {
		logger.Error(err.Error())
	}

	return CustomerRepositoryDB{dbClient: db}
}
