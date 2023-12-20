package infrastructure

import (
	"log"
	"stori/pkg/result"
	"stori/pkg/transaction"
	"strconv"
	"strings"
)

type TransactionDetails struct {
	Total             int
	TotalDebitAmount  float64
	TotalCreditAmount float64
}

type TransactionsCalculator struct{}

func (t *TransactionsCalculator) Calculate(col []transaction.Transaction) (result.Result, error) {
	_ = TotalBalance(col)
	return result.Result{}, nil
}

func TotalBalance(col []transaction.Transaction) float64 {
	var total float64
	monthlyInfo := map[int]TransactionDetails{}

	for i := 0; i < len(col); i++ {
		total += GetAmount(col[i].Transaction)
		month := GetMonth(col[i].Date)

		if _, ok := monthlyInfo[month]; !ok {
			monthlyInfo[month] = TransactionDetails{
				Total: 1,
			}
		} else {
			tmp := monthlyInfo[month]
			tmp.Total++
			monthlyInfo[month] = tmp
		}
	}

	log.Printf("%+v", monthlyInfo)

	return total
}

func GetMonth(date string) int {
	month := strings.Split(date, "/")[0]
	tmp, _ := strconv.ParseInt(month, 10, 64)
	return int(tmp)
}

func GetAmount(trx string) float64 {
	symbol := trx[0]
	tmpAmount := trx[1:]
	amount, _ := strconv.ParseFloat(tmpAmount, 64)

	if string(symbol) == "+" {
		return amount
	}
	return -amount
}
