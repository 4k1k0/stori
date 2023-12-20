package application

import (
	reader "stori/internal/reader/application"
	"stori/pkg/processor/domain"
	"stori/pkg/processor/infrastructure"
)

type ProcessorType string

const (
	ProcessorFile ProcessorType = "ProcessorFile"
)

func New(path string, pt ProcessorType) domain.Processor {
	return &infrastructure.ProcessorTransactions{
		File:   path,
		Reader: reader.New(),
	}
}
