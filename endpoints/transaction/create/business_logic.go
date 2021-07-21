package createtransaction

import (
	"net/http"

	"github.com/volatiletech/null/v8"
	"simpl.com/errors"
	. "simpl.com/loggers"
	. "simpl.com/repositories"
	"simpl.com/repositories/models"
	"simpl.com/utils"
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

	// credit limit check
	if utils.RoundUp(command.Amount, 2) > utils.RoundUp(user.CreditLimit.Float64, 2) - utils.RoundUp(user.DueAmount.Float64, 2) {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "amount exceeds credit_limit"

		return &transaction, businessError
	}

	totalAmount := utils.RoundUp(command.Amount, 2)
	discountAmount := utils.RoundUp(totalAmount * merchant.DiscountPercent.Float64 / 100, 2)
	paidAmount := utils.RoundUp(totalAmount - discountAmount, 2)

	Logger.Info(totalAmount)
	Logger.Info(discountAmount)
	Logger.Info(paidAmount)


	transaction.TotalAmount = totalAmount
	transaction.DiscountAmount =  null.Float64From(discountAmount)
	transaction.PaidAmount = paidAmount

	if err = Repositories.TransactionsRepository.PutTransaction(&transaction, nil); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &transaction, businessError
	}

	return &transaction, businessError
}