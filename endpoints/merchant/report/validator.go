package merchantreport

import (
	"strings"
)

func (merchantReportRequest *MerchantReportRequest) Validate() []string {
	var err []string

	if strings.TrimSpace(merchantReportRequest.ID) == "" {
		err = append(err, "required merchant ID in URL")
	}

	return err
}