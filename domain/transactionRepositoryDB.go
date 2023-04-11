package domain

import (
	"go-bank-api/errs"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryDB struct {
	client *sqlx.DB
}

func (tr TransactionRepositoryDB) CreateTransaction(t Transaction) (*Transaction, *errs.AppError) {
	sqlInsert := "INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)"

	result, err := tr.client.Exec(sqlInsert, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	t.Id = strconv.FormatInt(id, 10)
	return &t, nil

}

func NewTransactionRepositoryDB(dbClient *sqlx.DB) TransactionRepositoryDB {
	return TransactionRepositoryDB{client: dbClient}
}