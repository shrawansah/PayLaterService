package merchantreport

import (
	"net/http"
	"github.com/gorilla/mux"

	. "simpl.com/loggers"
)

func (merchantReportRequest *MerchantReportRequest) Decode(r *http.Request) error {
	merchantID := mux.Vars(r)["id"]
	merchantReportRequest.ID = merchantID

	Logger.Info("CreateMerchantRequest :: ", merchantReportRequest.ToString())
	return nil
}