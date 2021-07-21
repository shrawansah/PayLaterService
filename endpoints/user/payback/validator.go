package userpayback

func (paybackUserRequest *PaybackUserRequest) Validate() []string {
	var err []string

	if paybackUserRequest.UserID <= 0 {
		err = append(err, "invalid user_id")
	}
	if paybackUserRequest.Amount <= 0  {
		err = append(err, "invalid amount")
	}

	return err
}