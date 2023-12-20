package domain

import (
	"stori/pkg/result"
	"stori/pkg/transaction"
)

type Calculator interface {
	Calculate(col []transaction.Transaction) (result.Result, error)
}
