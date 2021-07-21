package createuser

import (
	"encoding/json"
)

type CreateUserRequest struct {
	Name 				string      `json:"name"`
	Email			 	string		`json:"email"`
	CreditLimit			int64		`json:"credit_limit"`
}

func (createUserRequest *CreateUserRequest) ToString() string {
	bytes, _ := json.Marshal(createUserRequest)
	return string(bytes)
}