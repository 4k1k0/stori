package transaction

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	PatterDate        = `^(1[0-2]|[1-9])/([1-9]|[1-2][0-9]|3[0-1])$`
	PatterTransaction = `^(\+[0-9]+(\.[0-9]+)?|\-[0-9]+(\.[0-9]+)?)$`
)

type Transaction struct {
	ID          int64  `json:"id"`
	Date        string `json:"date"`
	Transaction string `json:"transaction"`
}

func (t Transaction) IsCredit() bool {
	return string(t.Transaction[0]) == "+"
}

func (t Transaction) GetAmounts() (credit, debit float64) {
	tmpAmount := t.Transaction[1:]
	amount, _ := strconv.ParseFloat(tmpAmount, 64)

	if t.IsCredit() {
		credit = amount
		return
	}

	debit = -amount

	return
}

func (t Transaction) GetMonth() int {
	month := strings.Split(t.Date, "/")[0]
	tmp, _ := strconv.ParseInt(month, 10, 64)
	return int(tmp)
}

func (t Transaction) Type() string {
	if t.IsCredit() {
		return "Credit"
	}
	return "Debit"
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
