package application

import (
	reader "stori/internal/reader/application"
	"stori/pkg/processor/domain"
	"stori/pkg/processor/infrastructure"
)

func New(file string) domain.Processor {
	return &infrastructure.ProcessorTransactions{
		Reader:     reader.New(file),
		Calculator: nil,
	}
}
