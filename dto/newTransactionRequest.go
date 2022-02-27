package dto

type NewTransactionRequest struct {
	AccountId       int64   `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}
