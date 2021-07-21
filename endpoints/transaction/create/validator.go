package createtransaction

import (
	"simpl.com/utils"
)

func (createTransactionRequest *CreateTransactionRequest) Validate() []string {
	var err []string

	if createTransactionRequest.Amount <= 0  {
		err = append(err, "amount can not be less than or equal to zero")
	}
	if (!utils.CheckDecimalPlaces(2, createTransactionRequest.Amount)) {
		err = append(err, "amount can not have more than two decimal places")
	}
	if createTransactionRequest.UserID <= 0 {
		err = append(err, "invalid user_id")
	}
	if createTransactionRequest.MerchantID <= 0 {
		err = append(err, "invalid merchant_id")
	}

	return err
}