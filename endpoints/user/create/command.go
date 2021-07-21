package createuser

import (
	"strings"
	"encoding/json"

	. "simpl.com/loggers"
)

type CreateUserCommand struct {
	Name 				string      `json:"name"`
	Email			 	string		`json:"email"`
	CreditLimit			int64		`json:"credit_limit"`
}

func (createUserCommand *CreateUserCommand) ToString() string {
	bytes, _ := json.Marshal(createUserCommand)
	return string(bytes)
}


func (command *CreateUserCommand) BuildFromRequest(request *CreateUserRequest) {

	command.Name = strings.TrimSpace(request.Name)
	command.CreditLimit = request.CreditLimit
	command.Email = strings.TrimSpace(request.Email)

	Logger.Info("CreateUserCommand :: ", command.ToString())
}