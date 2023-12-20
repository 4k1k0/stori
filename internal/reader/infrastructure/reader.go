package infrastructure

import (
	"stori/pkg/transaction"

	"github.com/jszwec/csvutil"
)

type ReaderCSV struct{}

func (r *ReaderCSV) Read(data []byte) ([]transaction.Transaction, []byte, error) {
	var transactions transaction.Collection
	if err := csvutil.Unmarshal(data, &transactions); err != nil {
		return nil, nil, err
	}

	if err := transactions.Validate(); err != nil {
		return nil, nil, err
	}

	return transactions, data, nil
}
