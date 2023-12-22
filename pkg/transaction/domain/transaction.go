package domain

import (
	"database/sql"
	"stori/pkg/transaction"
)

type TransactionDataConnector interface {
	Save(db *sql.DB, trx transaction.Transaction) error
}
