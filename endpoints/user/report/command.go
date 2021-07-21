package merchantreport

import (
	"encoding/json"

	. "simpl.com/loggers"
)

type MerchantReportCommand struct {
	ID string
}

func (merchantReportCommand *MerchantReportCommand) ToString() string {
	bytes, _ := json.Marshal(merchantReportCommand)
	return string(bytes)
}


func (command *MerchantReportCommand) BuildFromRequest(request *MerchantReportRequest) {

	command.ID = request.ID
	Logger.Info("MerchantReportCommand :: ", command.ToString())
}