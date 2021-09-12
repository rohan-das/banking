package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rohan-das/banking/errs"
	"github.com/rohan-das/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (db CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		query := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = db.client.Query(query)
	} else {
		query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = db.client.Query(query, status)
	}

	if err != nil {
		logger.Error("Error while querying customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customers " + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (db CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	query := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"
	row := db.client.QueryRow(query, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		}
		logger.Error("Error while scanning customers " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3307)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return CustomerRepositoryDb{client: db}
}
