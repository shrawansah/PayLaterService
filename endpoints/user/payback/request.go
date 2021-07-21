package userpayback

import (
	"encoding/json"
)

type PaybackUserRequest struct {
	UserID 			int64       `json:"user_id"`
	Amount			int64		`json:"amount"`
}

func (paybackUserRequest *PaybackUserRequest) ToString() string {
	bytes, _ := json.Marshal(paybackUserRequest)
	return string(bytes)
}