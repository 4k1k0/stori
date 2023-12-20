package domain

import (
	"stori/pkg/transaction"
)

type Reader interface {
	Read() ([]transaction.Transaction, []byte, error)
}

type ReaderSlice interface {
	Read(data []byte) ([]transaction.Transaction, []byte, error)
}
