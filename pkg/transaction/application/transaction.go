package application

import (
	"stori/pkg/transaction/domain"
	"stori/pkg/transaction/infrastructure"
)

func New() domain.TransactionDataConnector {
	return &infrastructure.TransactionPostgres{}
}
