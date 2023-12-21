package application

import (
	"stori/pkg/processor/domain"
	"stori/pkg/processor/infrastructure"

	calculator "stori/internal/calculator/application"
	"stori/internal/config"
	reader "stori/internal/reader/application"
	sender "stori/internal/sender/application"
)

func New() domain.Processor {
	cfg := config.Config()
	return &infrastructure.ProcessorTransactions{
		Reader:     reader.New(cfg.FileName),
		Calculator: calculator.New(),
		Sender:     sender.New(cfg.Email),
	}
}
