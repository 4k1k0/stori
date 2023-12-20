package domain

import (
	"stori/pkg/transaction"
)

type Reader interface {
	Read() ([]transaction.Transaction, error)
}

type ReaderSlice interface {
	Read(data []byte) ([]transaction.Transaction, error)
}
