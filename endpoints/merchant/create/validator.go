package createmerchant

import (
	"strings"
)

func (createMerchantRequest *CreateMerchantRequest) Validate() []string {
	var err []string

	if strings.TrimSpace(createMerchantRequest.Name) == "" {
		err = append(err, "name can not be empty")
	}
	if createMerchantRequest.DiscountPercent < 0  || createMerchantRequest.DiscountPercent > 100 {
		err = append(err, "invalid discount_percent")
	}

	return err
}