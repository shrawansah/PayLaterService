package createmerchant

import (
	"net/http"

	"github.com/volatiletech/null/v8"
	. "simpl.com/loggers"
	. "simpl.com/repositories"
	"simpl.com/errors"
	"simpl.com/repositories/models"
)

func (command *CreateMerchantCommand) ExecuteBusinessLogic() (*models.Merchant, errors.BusinessLogicError) {
	
	merchant := models.Merchant {
		Name: command.Name,
		DiscountPercent: null.Float64From(float64(command.DiscountPercent)),
	}
	businessError := errors.BusinessLogicError{}
	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	merchants, err := Repositories.MerchantsRepository.GetMerchants("name = ? ", command.Name)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &merchant, businessError
	}
	if len(merchants) > 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "Merchant with same name already exists"

		return &merchant, businessError
	}

	if err := Repositories.MerchantsRepository.PutMerchant(&merchant, nil); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &merchant, businessError
	}

	return &merchant, businessError
}