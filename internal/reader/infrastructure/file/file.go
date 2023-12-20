package file

import "stori/pkg/transaction"

type ReaderFile struct {
	File string
}

func (r *ReaderFile) Read() ([]transaction.Transaction, error) {
	return nil, nil
}
