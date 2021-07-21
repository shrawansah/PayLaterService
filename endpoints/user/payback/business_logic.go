package userpayback

import (
	"fmt"
	"net/http"
	"context"

	"simpl.com/errors"
	. "simpl.com/loggers"
	. "simpl.com/repositories"
	"simpl.com/repositories/models"
	"simpl.com/databases"
)

func (command *PaybackUserCommand) ExecuteBusinessLogic() (*models.Payback, errors.BusinessLogicError) {
	
	payback := models.Payback {
		UserID: command.UserID,
		Amount: command.Amount,
	}
	businessError := errors.BusinessLogicError{}
	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	users, err := Repositories.UsersRepository.GetUsers("id = ? ", command.UserID)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &payback, businessError
	}
	if len(users) == 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "unknown user_id"

		return &payback, businessError
	}
	user := users[0]

	if user.DueAmount < command.Amount {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "can not pay more than " + fmt.Sprintf("%d", user.DueAmount)

		return &payback, businessError
	}
	sqlTxn, err := databases.GetConnection().BeginTx(context.Background(), nil)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &payback, businessError
	}

	user.DueAmount -= command.Amount
	if _, err := Repositories.UsersRepository.UpdateUser(user, sqlTxn); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"
		sqlTxn.Rollback()
		return &payback, businessError
	}

	if err := Repositories.PaybacksRepository.PutPayback(&payback, sqlTxn); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"
		sqlTxn.Rollback()
		return &payback, businessError
	}

	sqlTxn.Commit()
	return &payback, businessError
}