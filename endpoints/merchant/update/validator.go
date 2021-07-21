package updatemerchant

func (updateMerchantRequest *UpdateMerchantRequest) Validate() []string {
	var err []string

	if updateMerchantRequest.ID <= 0 {
		err = append(err, "invalid id")
	}

	if updateMerchantRequest.DiscountPercent < 0  || updateMerchantRequest.DiscountPercent > 100 {
		err = append(err, "invalid discount_percent")
	}

	return err
}