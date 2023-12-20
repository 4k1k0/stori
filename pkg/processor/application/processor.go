package application

import (
	"stori/pkg/processor/domain"
	"stori/pkg/processor/infrastructure"

	calculator "stori/internal/calculator/application"
	reader "stori/internal/reader/application"
	sender "stori/internal/sender/application"
)

func New(file, email string) domain.Processor {
	return &infrastructure.ProcessorTransactions{
		Reader:     reader.New(file),
		Calculator: calculator.New(),
		Sender:     sender.New(email),
	}
}
