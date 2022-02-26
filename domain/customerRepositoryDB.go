package domain

import (
	"database/sql"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/rtpa25/banking/errs"
	"github.com/rtpa25/banking/logger"
	"github.com/spf13/viper"
)

type CustomerRepositoryDB struct {
	dbClient *sql.DB
}

func (d CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSQL string
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSQL = `SELECT id, name, city, zipcode, date_of_birth, status from customers`
		rows, err = d.dbClient.Query(findAllSQL)
	} else {
		findAllSQL = `SELECT id, name, city, zipcode, date_of_birth, status from customers WHERE status = $1`
		if status == "active" {
			rows, err = d.dbClient.Query(findAllSQL, 1)
		} else if status == "inactive" {
			rows, err = d.dbClient.Query(findAllSQL, 0)
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

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customer" + err.Error())
			return nil, errs.NewUnexpectedError("unexpected database error")
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDB) ById(id string) (*Customer, *errs.AppError) {
	findUserSQL := `SELECT id, name, city, zipcode, date_of_birth, status from customers WHERE id = $1`
	intId, err := strconv.Atoi(id)
	if err != nil {
		logger.Error(err.Error())
	}
	row := d.dbClient.QueryRow(findUserSQL, intId) //only return one row
	var c Customer
	err = row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
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
	db, err := sql.Open("postgres", connStr.(string))
	if err != nil {
		logger.Error(err.Error())
	}

	return CustomerRepositoryDB{dbClient: db}
}
