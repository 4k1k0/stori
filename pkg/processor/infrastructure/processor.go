package infrastructure

import (
	"log"
	"stori/pkg/result"

	calculator "stori/internal/calculator/domain"
	reader "stori/internal/reader/domain"
	sender "stori/internal/sender/domain"
)

type ProcessorTransactions struct {
	Calculator calculator.Calculator
	Reader     reader.Reader
	Sender     sender.Sender
}

func (p *ProcessorTransactions) Process() (result.Result, error) {
	data, info, err := p.Reader.Read()
	if err != nil {
		return result.Result{}, err
	}

	res, err := p.Calculator.Calculate(data)
	if err != nil {
		return result.Result{}, err
	}

	err = p.Sender.Send(res, info)
	if err != nil {
		log.Println("error sending email")
		log.Println(err)
	}

	return res, nil
}
