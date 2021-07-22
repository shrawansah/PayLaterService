package createtransaction

import (
	"context"
	"net/http"

	"simpl.com/databases"
	"simpl.com/errors"
	. "simpl.com/loggers"
	. "simpl.com/repositories"
	"simpl.com/repositories/models"
)

func (command *CreateTransactionCommand) ExecuteBusinessLogic() (*models.Transaction, errors.BusinessLogicError) {
	
	transaction := models.Transaction {
		UserID: 	command.UserID,
		MerchantID: command.MerchantID,
		TotalAmount: command.Amount,
	}
	merchant := models.Merchant{}
	user := models.User{}

	businessError := errors.BusinessLogicError{}
	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	// merchant check
	merchants, err := Repositories.MerchantsRepository.GetMerchants("id = ? ", command.MerchantID)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &transaction, businessError
	}
	if len(merchants) == 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "merchant_id not known"

		return &transaction, businessError
	}
	merchant = *merchants[0]

	// user check
	users, err := Repositories.UsersRepository.GetUsers("id = ? ", command.UserID)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &transaction, businessError
	}
	if len(users) == 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "user_id not known"

		return &transaction, businessError
	}
	user = *users[0]

	Logger.Info("total_amount :: ", command.Amount)
	Logger.Info("credit_limit :: ", user.CreditLimit)
	Logger.Info("due_amount :: ", user.DueAmount)


	// credit limit check
	if command.Amount > user.CreditLimit - user.DueAmount {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "amount exceeds credit_limit"

		return &transaction, businessError
	}

	sqlTxn, err := databases.GetConnection().BeginTx(context.Background(), nil)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &transaction, businessError
	}


	totalAmount := command.Amount
	discountAmount := int64(float64(totalAmount) * merchant.DiscountPercent.Float64 / 100)
	paidAmount := totalAmount - discountAmount

	transaction.TotalAmount = totalAmount
	transaction.DiscountAmount =  discountAmount
	transaction.PaidAmount = paidAmount

	user.DueAmount += totalAmount
	if _, err := Repositories.UsersRepository.UpdateUser(&user, sqlTxn); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"
		sqlTxn.Rollback()
		return &transaction, businessError
	}

	if err = Repositories.TransactionsRepository.PutTransaction(&transaction, sqlTxn); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"
		sqlTxn.Rollback()
		return &transaction, businessError
	}

	sqlTxn.Commit()
	return &transaction, businessError
}