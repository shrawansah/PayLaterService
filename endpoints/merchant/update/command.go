package updatemerchant

import (
	"encoding/json"

	. "simpl.com/loggers"
)

type UpdateMerchantCommand struct {
	ID 					int			`json:"id"`
	DiscountPercent 	float64		`json:"discount_percent"`
}

func (updateMerchantCommand *UpdateMerchantCommand) ToString() string {
	bytes, _ := json.Marshal(updateMerchantCommand)
	return string(bytes)
}


func (command *UpdateMerchantCommand) BuildFromRequest(request *UpdateMerchantRequest) {
	command.ID = request.ID
	command.DiscountPercent = request.DiscountPercent
	Logger.Info("UpdateMerchantCommand :: ", command.ToString())
}