package createuser

import (
	"net/http"

	"github.com/volatiletech/null/v8"
	. "simpl.com/loggers"
	. "simpl.com/repositories"
	"simpl.com/errors"
	"simpl.com/repositories/models"
)

func (command *CreateUserCommand) ExecuteBusinessLogic() (*models.User, errors.BusinessLogicError) {
	
	user := models.User {
		Name: command.Name,
		CreditLimit: null.Float64From(command.CreditLimit),
		EmailID: command.Email,
		DueAmount: null.Float64From(float64(0)),
	}
	businessError := errors.BusinessLogicError{}
	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	users, err := Repositories.UsersRepository.GetUsers("email_id = ? ", command.Email)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &user, businessError
	}
	if len(users) > 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "Merchant with same name already exists"

		return &user, businessError
	}

	if err := Repositories.UsersRepository.PutUser(&user, nil); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &user, businessError
	}

	return &user, businessError
}