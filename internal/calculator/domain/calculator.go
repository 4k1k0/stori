package domain

import "stori/pkg/transaction"

type Calculator interface {
	Calculate([]transaction.Transaction) (interface{}, error)
}
