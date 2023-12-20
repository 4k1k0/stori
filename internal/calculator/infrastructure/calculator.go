package infrastructure

import (
	"log"
	"stori/pkg/result"
	"stori/pkg/transaction"
)

type Details struct {
	Total      float64
	Collection CollectionTrxDetails
}

type CollectionTrxDetails map[int]TransactionDetails

type TransactionDetails struct {
	Total             int
	TotalDebitAmount  float64
	TotalCreditAmount float64
}

type TransactionsCalculator struct{}

func (t *TransactionsCalculator) Calculate(col []transaction.Transaction) (result.Result, error) {
	tmp := Process(col)
	res := GetResult(tmp)

	return res, nil
}

func Process(col []transaction.Transaction) Details {
	var total float64
	monthlyInfo := CollectionTrxDetails{}

	for i := 0; i < len(col); i++ {
		c, d := col[i].GetAmounts()
		total += c
		total += d
		month := col[i].GetMonth()

		if _, ok := monthlyInfo[month]; !ok {
			monthlyInfo[month] = TransactionDetails{
				Total:             1,
				TotalDebitAmount:  d,
				TotalCreditAmount: c,
			}
		} else {
			tmp := monthlyInfo[month]
			tmp.Total++
			tmp.TotalCreditAmount += c
			tmp.TotalDebitAmount += d
			monthlyInfo[month] = tmp
		}
	}

	return Details{
		Total:      total,
		Collection: monthlyInfo,
	}
}

func GetResult(res Details) result.Result {
	transactions := make([]result.TransactionDetails, len(res.Collection))
	var i int

	for k, v := range res.Collection {
		log.Println(k, v)
		transactions[i] = result.TransactionDetails{
			Month:                k,
			NumberOfTransactions: v.Total,
			AverageDebitAmount:   v.TotalDebitAmount / float64(v.Total),
			AverageCreditAmount:  v.TotalCreditAmount / float64(v.Total),
		}
		i++
	}

	return result.Result{
		TotalBalance: res.Total,
		Transactions: transactions,
	}
}
