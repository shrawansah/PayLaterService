package createtransaction

import (
	"encoding/json"
)

type CreateTransactionRequest struct {
	UserID 		uint64      `json:"user_id"`
	MerchantID 	uint64		`json:"merchant_id"`
	Amount		int64		`json:"amount"`
}

func (createTransactionRequest *CreateTransactionRequest) ToString() string {
	bytes, _ := json.Marshal(createTransactionRequest)
	return string(bytes)
}