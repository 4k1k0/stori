package infrastructure

import (
	"log"
	"stori/pkg/result"
	"stori/pkg/transaction"
)

type TransactionsCalculator struct{}

func (t *TransactionsCalculator) Calculate([]transaction.Transaction) (result.Result, error) {
	log.Println("hello from my calculator...")
	return result.Result{}, nil
}
