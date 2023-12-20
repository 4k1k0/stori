package infrastructure

import (
	"fmt"

	calculator "stori/internal/calculator/domain"
	reader "stori/internal/reader/domain"
	"stori/pkg/result"
)

type ProcessorTransactions struct {
	Calculator calculator.Calculator
	Reader     reader.Reader
}

func (p *ProcessorTransactions) Process() (result.Result, error) {
	data, err := p.Reader.Read()
	if err != nil {
		return result.Result{}, err
	}

	fmt.Println(data)

	return result.Result{}, nil
}
