package infrastructure

import (
	"database/sql"
	"stori/pkg/transaction"
	"time"
)

const (
	queryCreateTable = `CREATE TABLE IF NOT EXISTS "public.Transactions" (
	"id" serial NOT NULL,
	"filename" varchar(100) NOT NULL,
	"date" varchar(5) NOT NULL,
	"created_at" TIMESTAMP NOT NULL,
	"transaction" varchar(20) NOT NULL,
	"transaction_type" varchar(10) NOT NULL,
	CONSTRAINT "Transactions_pk" PRIMARY KEY ("id")
) WITH (
  OIDS=FALSE
);`
	querySaveTransaction = `INSERT INTO "public.Transactions" ("filename", "date", "created_at", "transaction", "transaction_type") VALUES ($1, $2, $3, $4, $5)`
)

type TransactionPostgres struct{}

func createTable(db *sql.DB) error {
	_, err := db.Exec(queryCreateTable)
	return err
}

func (*TransactionPostgres) Save(db *sql.DB, trx transaction.Transaction) error {
	if err := createTable(db); err != nil {
		return err
	}

	tn := time.Now()

	_, err := db.Exec(querySaveTransaction, "filename", trx.Date, tn, trx.Transaction, trx.Type())

	return err
}
