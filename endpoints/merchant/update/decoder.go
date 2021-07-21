package updatemerchant

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "simpl.com/loggers"
)

func (updateMerchantRequest *UpdateMerchantRequest) Decode(r *http.Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error(err)
		return err
	}

	if err := json.Unmarshal(reqBody, &updateMerchantRequest); err != nil {
		Logger.Error(err)
		return err
	}
	Logger.Info("UpdateMerchantRequest :: ", updateMerchantRequest.ToString())
	return nil
}