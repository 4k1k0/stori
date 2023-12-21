package infrastructure

import (
	"log"
	"stori/pkg/result"
	"stori/pkg/transaction"

	calculator "stori/internal/calculator/domain"
	"stori/internal/config"
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

	for i := 0; i < len(data); i++ {
		err = transaction.SaveTransaction(config.Config().Database, data[i])
		if err != nil {
			break
		}
	}

	if err != nil {
		log.Println("There was an error saving the data")
		log.Println(err)
	}

	return res, nil
}
