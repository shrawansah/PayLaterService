package updatemerchant

import (
	"simpl.com/utils"
)

func (updateMerchantRequest *UpdateMerchantRequest) Validate() []string {
	var err []string

	if updateMerchantRequest.ID <= 0 {
		err = append(err, "invalid id")
	}

	if updateMerchantRequest.DiscountPercent < 0  || updateMerchantRequest.DiscountPercent > 100 {
		err = append(err, "invalid discount_percent")
	}

	if (!utils.CheckDecimalPlaces(2, updateMerchantRequest.DiscountPercent)) {
		err = append(err, "discount_percent can not have more than two decimal places")
	}

	return err
}