package userpayback

import (
	"encoding/json"

	. "simpl.com/loggers"
)

type PaybackUserCommand struct {
	UserID 			int64       `json:"user_id"`
	Amount			int64		`json:"amount"`
}

func (createUserCommand *PaybackUserCommand) ToString() string {
	bytes, _ := json.Marshal(createUserCommand)
	return string(bytes)
}


func (command *PaybackUserCommand) BuildFromRequest(request *PaybackUserRequest) {

	command.UserID = request.UserID
	command.Amount = request.Amount

	Logger.Info("PaybackUserCommand :: ", command.ToString())
}