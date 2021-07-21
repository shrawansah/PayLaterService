package userreport

import (
	"net/http"

	. "simpl.com/loggers"
	"simpl.com/repositories"
	"simpl.com/errors"
)

func (command *UserReportCommand) ExecuteBusinessLogic() (repositories.UserStatistics, errors.BusinessLogicError) {
	
	businessError := errors.BusinessLogicError{}
	var propagator repositories.UserStatistics

	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	propagator, err := repositories.Repositories.UsersRepository.GetAllStats(command.ID)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return propagator, businessError
	}

	return propagator, businessError
}