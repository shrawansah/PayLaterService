package createtransaction

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "simpl.com/loggers"
)

func (createTransactionRequest *CreateTransactionRequest) Decode(r *http.Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error(err)
		return err
	}

	if err := json.Unmarshal(reqBody, &createTransactionRequest); err != nil {
		Logger.Error(err)
		return err
	}
	Logger.Info("CreateTransactionRequest :: ", createTransactionRequest.ToString())
	return nil
}