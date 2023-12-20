package infrastructure

import (
	"fmt"
	"log"

	calculator "stori/internal/calculator/domain"
	reader "stori/internal/reader/domain"
)

type ProcessorTransactions struct {
	Calculator calculator.Calculator
	File       string
	Reader     reader.Reader
}

func (p *ProcessorTransactions) Process() (interface{}, error) {
	log.Println("Opening file", p.File)

	data, err := p.Reader.Read()
	if err != nil {
		return nil, err
	}

	fmt.Println(data)

	return nil, nil
}
