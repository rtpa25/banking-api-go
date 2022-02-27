package dto

type NewTransactionResponse struct {
	Id                  int64   `json:"id"`
	UpdatedBankBallance float64 `json:"updated_ballance"`
}
