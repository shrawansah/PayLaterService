package createtransaction

import (
	"encoding/json"

	. "simpl.com/loggers"
	"simpl.com/utils"
)

type CreateTransactionCommand struct {
	UserID 		uint64      `json:"user_id"`
	MerchantID 	uint64		`json:"merchant_id"`
	Amount		float64		`json:"amount"`
}

func (createMerchantCommand *CreateTransactionCommand) ToString() string {
	bytes, _ := json.Marshal(createMerchantCommand)
	return string(bytes)
}


func (command *CreateTransactionCommand) BuildFromRequest(request *CreateTransactionRequest) {

	command.UserID = request.UserID
	command.MerchantID = request.MerchantID
	command.Amount = utils.RoundUp(request.Amount, 2)
	Logger.Info("CreateTransactionCommand :: ", command.ToString())
}