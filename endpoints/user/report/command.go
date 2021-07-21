package userreport

import (
	"encoding/json"

	. "simpl.com/loggers"
)

type UserReportCommand struct {
	ID string
}

func (userReportCommand *UserReportCommand) ToString() string {
	bytes, _ := json.Marshal(userReportCommand)
	return string(bytes)
}


func (command *UserReportCommand) BuildFromRequest(request *UserReportRequest) {

	command.ID = request.ID
	if request.ID == "all" {
		command.ID = ""
	}
	Logger.Info("UserReportCommand :: ", command.ToString())
}