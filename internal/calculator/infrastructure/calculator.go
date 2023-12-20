package infrastructure

import (
	"sort"
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
	TotalDebit        int
	TotalCredit       int
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
		var ct, dt int
		if col[i].IsCredit() {
			ct++
		} else {
			dt++
		}
		total += c
		total += d
		month := col[i].GetMonth()

		if _, ok := monthlyInfo[month]; !ok {
			monthlyInfo[month] = TransactionDetails{
				Total:             1,
				TotalDebitAmount:  d,
				TotalCreditAmount: c,
				TotalDebit:        dt,
				TotalCredit:       ct,
			}
		} else {
			tmp := monthlyInfo[month]
			tmp.Total++
			tmp.TotalCreditAmount += c
			tmp.TotalDebitAmount += d
			tmp.TotalDebit += dt
			tmp.TotalCredit += ct
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
		var averageDA, averageCA float64
		if v.TotalDebit > 0 {
			tmp := v.TotalDebitAmount / float64(v.TotalDebit)
			averageDA = tmp
		}

		if v.TotalCredit > 0 {
			tmp := v.TotalCreditAmount / float64(v.TotalCredit)
			averageCA = tmp
		}

		transactions[i] = result.TransactionDetails{
			Month:                k,
			MonthName:            months[k],
			NumberOfTransactions: v.Total,
			AverageDebitAmount:   averageDA,
			AverageCreditAmount:  averageCA,
		}
		i++
	}

	sort.SliceStable(transactions, func(i, j int) bool {
		return transactions[i].Month < transactions[j].Month
	})

	return result.Result{
		TotalBalance: res.Total,
		Transactions: transactions,
	}
}
