package userreport

import (
	"net/http"
	"github.com/gorilla/mux"

	. "simpl.com/loggers"
)

func (userReportRequest *UserReportRequest) Decode(r *http.Request) error {
	merchantID := mux.Vars(r)["id"]
	userReportRequest.ID = merchantID

	Logger.Info("CreateMerchantRequest :: ", userReportRequest.ToString())
	return nil
}