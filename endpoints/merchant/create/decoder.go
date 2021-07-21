package createmerchant

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "simpl.com/loggers"
)

func (createMerchantRequest *CreateMerchantRequest) Decode(r *http.Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error(err)
		return err
	}

	if err := json.Unmarshal(reqBody, &createMerchantRequest); err != nil {
		Logger.Error(err)
		return err
	}
	Logger.Info("CreateMerchantRequest :: ", createMerchantRequest.ToString())
	return nil
}