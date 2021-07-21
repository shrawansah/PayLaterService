package merchantreport

import (
	"net/http"

	. "simpl.com/loggers"
	repo "simpl.com/repositories"
	"simpl.com/errors"
)

func (command *MerchantReportCommand) ExecuteBusinessLogic() (*repo.StatisticsPropagator, errors.BusinessLogicError) {
	
	businessError := errors.BusinessLogicError{}
	propagator := repo.StatisticsPropagator{}

	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	merchants, err := repo.Repositories.MerchantsRepository.GetMerchants("id = ? ", command.ID)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &propagator, businessError
	}
	if len(merchants) == 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "Merchant does not exist"

		return &propagator, businessError
	}

	propagator, err = repo.Repositories.MerchantsRepository.GetAllStats(command.ID)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &propagator, businessError
	}

	return &propagator, businessError
}