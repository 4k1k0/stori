package application

import (
	"stori/pkg/processor/domain"
	"stori/pkg/processor/infrastructure"
)

func New(path string) domain.Processor {
	return &infrastructure.ProcessorTransactions{}
}
