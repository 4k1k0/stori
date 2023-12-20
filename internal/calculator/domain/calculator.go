package domain

import (
	"stori/pkg/result"
	"stori/pkg/transaction"
)

type Calculator interface {
	Calculate([]transaction.Transaction) (result.Result, error)
}
