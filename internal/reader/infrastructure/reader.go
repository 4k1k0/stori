package infrastructure

import (
	"log"
	"stori/pkg/transaction"
)

type ReaderCSV struct{}

func (r *ReaderCSV) Read() ([]transaction.Transaction, error) {
	log.Println("read method from ReaderCSV")
	return nil, nil
}
