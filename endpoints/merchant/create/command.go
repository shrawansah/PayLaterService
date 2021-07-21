package createmerchant

import (
	"strings"
	"encoding/json"

	. "simpl.com/loggers"
)

type CreateMerchantCommand struct {
	Name 				string      `json:"name"`
	DiscountPercent 	float64		`json:"discount_percent"`
}

func (createMerchantCommand *CreateMerchantCommand) ToString() string {
	bytes, _ := json.Marshal(createMerchantCommand)
	return string(bytes)
}


func (command *CreateMerchantCommand) BuildFromRequest(request *CreateMerchantRequest) {

	command.Name = strings.TrimSpace(request.Name)
	command.DiscountPercent = request.DiscountPercent

	Logger.Info("CreateMerchantCommand :: ", command.ToString())
}