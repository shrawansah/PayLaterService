package createtransaction

func (createTransactionRequest *CreateTransactionRequest) Validate() []string {
	var err []string

	if createTransactionRequest.Amount < 10  {
		err = append(err, "amount can not be less than or equal to 10")
	}
	if createTransactionRequest.UserID <= 0 {
		err = append(err, "invalid user_id")
	}
	if createTransactionRequest.MerchantID <= 0 {
		err = append(err, "invalid merchant_id")
	}

	return err
}