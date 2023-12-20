package infrastructure

import (
	"log"
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

	log.Println(data)

	res, err := p.Calculator.Calculate(data)
	if err != nil {
		return result.Result{}, err
	}
	log.Println(res)

	return result.Result{}, nil
}
