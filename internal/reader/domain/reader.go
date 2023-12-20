package domain

import (
	"stori/pkg/transaction"
)

type Reader interface {
	Read() ([]transaction.Transaction, error)
}
