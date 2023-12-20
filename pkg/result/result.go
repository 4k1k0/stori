package result

type Result struct {
	TotalBalance float64            `json:"totalBalance"`
	Transactions TransactionDetails `json:"transactions"`
	DebitAmount  float64            `json:"debitAmount"`
	CreditAmount float64            `json:"creditAmount"`
}

type TransactionDetails struct {
	Month int     `json:"month"`
	Total float64 `json:"total"`
}
