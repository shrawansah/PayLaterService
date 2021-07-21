package updatemerchant

import (
	"net/http"

	"github.com/volatiletech/null/v8"
	. "simpl.com/loggers"
	. "simpl.com/repositories"
	"simpl.com/errors"
	"simpl.com/repositories/models"
)

func (command *UpdateMerchantCommand) ExecuteBusinessLogic() (*models.Merchant, errors.BusinessLogicError) {
	
	var merchant *models.Merchant

	businessError := errors.BusinessLogicError{}
	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	merchants, err := Repositories.MerchantsRepository.GetMerchants("id = ? ", command.ID)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return merchant, businessError
	}
	if len(merchants) == 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "Merchant not known"

		return merchant, businessError
	}

	merchant = merchants[0]
	merchant.DiscountPercent = null.Float64From(command.DiscountPercent)

	if _, err := Repositories.MerchantsRepository.UpdateMerchant(merchant, nil); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return merchant, businessError
	}

	return merchant, businessError
}