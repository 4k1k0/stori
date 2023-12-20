package domain

import (
	"io"
	"stori/pkg/transaction"
)

type Reader interface {
	Read() ([]transaction.Transaction, error)
}

type ReaderIO interface {
	Read(r io.Reader) ([]transaction.Transaction, error)
}
