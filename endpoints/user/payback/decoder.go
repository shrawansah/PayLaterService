package userpayback

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "simpl.com/loggers"
)

func (paybackUserRequest *PaybackUserRequest) Decode(r *http.Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error(err)
		return err
	}

	if err := json.Unmarshal(reqBody, &paybackUserRequest); err != nil {
		Logger.Error(err)
		return err
	}
	Logger.Info("PaybackUserRequest :: ", paybackUserRequest.ToString())
	return nil
}