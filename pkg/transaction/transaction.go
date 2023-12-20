package transaction

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	PatterDate        = `^(1[0-2]|[1-9])/([1-9]|[1-2][0-9]|3[0-1])$`
	PatterTransaction = `^(\+[0-9]+(\.[0-9]+)?|\-[0-9]+(\.[0-9]+)?)$`
)

type Collection []Transaction

type Transaction struct {
	ID          int64  `json:"id"`
	Date        string `json:"date"`
	Transaction string `json:"transaction"`
}

func (t Transaction) Validate() error {
	err := t.validateDate()
	if err != nil {
		return err
	}

	err = t.validateTransaction()
	if err != nil {
		return err
	}

	return nil
}

func (t Transaction) validateData(pattern, data string) error {
	re := regexp.MustCompile(pattern)
	if re.MatchString(data) {
		return nil
	}
	msg := fmt.Sprintf("invalid data %q", data)
	return errors.New(msg)
}

func (t Transaction) validateDate() error {
	return t.validateData(PatterDate, t.Date)
}

func (t Transaction) validateTransaction() error {
	return t.validateData(PatterTransaction, t.Transaction)
}

func (c Collection) Validate() error {
	if len(c) == 0 {
		return nil
	}

	var err error

	for i := 0; i < len(c); i++ {
		err = c[i].Validate()
		if err != nil {
			break
		}
	}

	return err
}
