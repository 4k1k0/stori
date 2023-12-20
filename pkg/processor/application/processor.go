package application

import (
	"stori/pkg/processor/domain"
	"stori/pkg/processor/infrastructure"

	calculator "stori/internal/calculator/application"
	reader "stori/internal/reader/application"
)

func New(file string) domain.Processor {
	return &infrastructure.ProcessorTransactions{
		Reader:     reader.New(file),
		Calculator: calculator.New(),
	}
}
