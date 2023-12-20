package result

type Result struct {
	TotalBalance float64              `json:"totalBalance"`
	Transactions []TransactionDetails `json:"transactions"`
}

type TransactionDetails struct {
	Month                int     `json:"month"`
	NumberOfTransactions int     `json:"numberOfTransactions"`
	AverageDebitAmount   float64 `json:"averageDebitAmount"`
	AverageCreditAmount  float64 `json:"averageCreditAmount"`
}
