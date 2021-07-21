package updatemerchant

import (
	"encoding/json"
)

type UpdateMerchantRequest struct {
	ID 					int			`json:"id"`
	DiscountPercent 	float64		`json:"discount_percent"`
}

func (updateMerchantRequest *UpdateMerchantRequest) ToString() string {
	bytes, _ := json.Marshal(updateMerchantRequest)
	return string(bytes)
}