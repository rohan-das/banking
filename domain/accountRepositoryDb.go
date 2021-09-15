package domain

import (
	"database/sql"
	"strconv"

	"github.com/rohan-das/banking/errs"
	"github.com/rohan-das/banking/logger"
)

type AccountRepositoryDb struct {
	client *sql.DB
}

func (db AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	query := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"

	result, err := db.client.Exec(query, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}

func NewAccountRepositoryDb(client *sql.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}
}