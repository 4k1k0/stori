package infrastructure

import (
	"log"
	"stori/pkg/transaction"

	"github.com/jszwec/csvutil"
)

type ReaderCSV struct{}

func (r *ReaderCSV) Read(data []byte) ([]transaction.Transaction, error) {
	log.Println("read method from ReaderCSV")

	var transactions []transaction.Transaction
	err := csvutil.Unmarshal(data, &transactions)
	if err != nil {
		return nil, err
	}

	return transactions, nil
}
