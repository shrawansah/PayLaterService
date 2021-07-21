package createmerchant

import (
	"encoding/json"
)

type CreateMerchantRequest struct {
	Name 				string      `json:"name"`
	DiscountPercent 	float64		`json:"discount_percent"`
}

func (createMerchantRequest *CreateMerchantRequest) ToString() string {
	bytes, _ := json.Marshal(createMerchantRequest)
	return string(bytes)
}