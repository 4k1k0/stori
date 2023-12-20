package result

type Result struct {
	TotalBalance float64              `json:"totalBalance"`
	Transactions []TransactionDetails `json:"transactions"`
}

type TransactionDetails struct {
	Month                int     `json:"month,omitempty"`
	MonthName            string  `json:"monthName,omitempty"`
	NumberOfTransactions int     `json:"numberOfTransactions,omitempty"`
	AverageDebitAmount   float64 `json:"averageDebitAmount,omitempty"`
	AverageCreditAmount  float64 `json:"averageCreditAmount,omitempty"`
	TotalDebit           int     `json:"totalDebit,omitempty"`
	TotalCredit          int     `json:"totalCredit,omitempty"`
}

type MonthSorter []TransactionDetails

func (a MonthSorter) Len() int           { return len(a) }
func (a MonthSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MonthSorter) Less(i, j int) bool { return a[i].MonthName < a[j].MonthName }
