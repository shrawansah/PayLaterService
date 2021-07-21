package createuser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	. "simpl.com/loggers"
)

func (createUserRequest *CreateUserRequest) Decode(r *http.Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		Logger.Error(err)
		return err
	}

	if err := json.Unmarshal(reqBody, &createUserRequest); err != nil {
		Logger.Error(err)
		return err
	}
	Logger.Info("CreateUserRequest :: ", createUserRequest.ToString())
	return nil
}