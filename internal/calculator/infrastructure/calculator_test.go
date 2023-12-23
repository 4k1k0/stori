package infrastructure

import (
	"stori/pkg/transaction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionsCalculator_Calculate(t *testing.T) {
	input := []transaction.Transaction{
		{ID: 1, Date: "1/11", Transaction: "+10"},
		{ID: 2, Date: "1/12", Transaction: "+10"},
		{ID: 3, Date: "1/13", Transaction: "-10"},
		{ID: 3, Date: "1/14", Transaction: "-10"},
	}
	res, err := (&TransactionsCalculator{}).Calculate(input)
	assert.Nil(t, err)
	assert.Equal(t, float64(0), res.TotalBalance)
	assert.Equal(t, 1, len(res.Transactions))
	assert.Equal(t, 1, res.Transactions[0].Month)
	assert.Equal(t, "January", res.Transactions[0].MonthName)
	assert.Equal(t, 4, res.Transactions[0].NumberOfTransactions)
	assert.Equal(t, float64(-10), res.Transactions[0].AverageDebitAmount)
	assert.Equal(t, float64(10), res.Transactions[0].AverageCreditAmount)
	assert.Equal(t, 0, res.Transactions[0].TotalDebit)
	assert.Equal(t, 0, res.Transactions[0].TotalCredit)
}

func TestProcess(t *testing.T) {
	input := []transaction.Transaction{
		{ID: 1, Date: "1/11", Transaction: "+10"},
		{ID: 2, Date: "1/12", Transaction: "+10"},
		{ID: 3, Date: "1/13", Transaction: "-10"},
		{ID: 3, Date: "1/14", Transaction: "-10"},
	}

	res := Process(input)
	assert.Equal(t, float64(0), res.Total)
	assert.Equal(t, 1, len(res.Collection))
	v, ok := res.Collection[1]
	assert.True(t, ok)
	assert.Equal(t, 4, v.Total)
	assert.Equal(t, 2, v.TotalCredit)
	assert.Equal(t, 2, v.TotalDebit)
	assert.Equal(t, float64(-20), v.TotalDebitAmount)
	assert.Equal(t, float64(20), v.TotalCreditAmount)
}

func TestGetResult(t *testing.T) {
	t.Run("without sorting", func(t *testing.T) {
		input := Details{
			Total: 1,
			Collection: CollectionTrxDetails{
				1: {
					Total:             1,
					TotalDebitAmount:  10,
					TotalCreditAmount: 0,
					TotalDebit:        10,
					TotalCredit:       0,
				},
			},
		}
		res := GetResult(input)
		assert.Equal(t, float64(1), res.TotalBalance)
		assert.Equal(t, 1, len(res.Transactions))
		assert.Equal(t, 1, res.Transactions[0].Month)
		assert.Equal(t, "January", res.Transactions[0].MonthName)
		assert.Equal(t, 1, res.Transactions[0].NumberOfTransactions)
		assert.Equal(t, float64(1), res.Transactions[0].AverageDebitAmount)
		assert.Equal(t, float64(0), res.Transactions[0].AverageCreditAmount)
	})

	t.Run("sorting", func(t *testing.T) {
		input := Details{
			Total: 1,
			Collection: CollectionTrxDetails{
				1: {},
				2: {},
				9: {},
			},
		}
		res := GetResult(input)
		assert.Equal(t, 3, len(res.Transactions))
		assert.Equal(t, "January", res.Transactions[0].MonthName)
		assert.Equal(t, "February", res.Transactions[1].MonthName)
		assert.Equal(t, "September", res.Transactions[2].MonthName)
	})
}
