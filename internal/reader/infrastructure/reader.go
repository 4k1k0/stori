package infrastructure

import (
	"log"
	"stori/pkg/transaction"

	"github.com/jszwec/csvutil"
)

type ReaderCSV struct{}

func (r *ReaderCSV) Read(data []byte) ([]transaction.Transaction, error) {
	log.Println("read method from ReaderCSV")

	var transactions transaction.Collection
	err := csvutil.Unmarshal(data, &transactions)
	if err != nil {
		return nil, err
	}

	if err = transactions.Validate(); err != nil {
		return nil, err
	}

	return transactions, nil
}
