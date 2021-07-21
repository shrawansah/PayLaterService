package merchantreport

import (
	"encoding/json"
)

type MerchantReportRequest struct {
	ID 		string
}

func (merchantReportRequest *MerchantReportRequest) ToString() string {
	bytes, _ := json.Marshal(merchantReportRequest)
	return string(bytes)
}