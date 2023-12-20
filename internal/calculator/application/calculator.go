package application

import (
	"stori/internal/calculator/domain"
	"stori/internal/calculator/infrastructure"
)

func New() domain.Calculator {
	return &infrastructure.TransactionsCalculator{}
}
