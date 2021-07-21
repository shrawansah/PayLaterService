package requests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/volatiletech/null/v8"
	. "simpl.com/loggers"
	. "simpl.com/repositories"
	"simpl.com/errors"
	"simpl.com/repositories/models"
)

type CreateMerchantRequest struct {
	Name 				string      `json:"name"`
	DiscountPercent 	float64		`json:"discount_percent"`
}

type CreateMerchantCommand struct {
	Name 				string      `json:"name"`
	DiscountPercent 	float64		`json:"discount_percent"`
}

func (createMerchantRequest *CreateMerchantRequest) ToString() string {
	bytes, _ := json.Marshal(createMerchantRequest)
	return string(bytes)
}
func (createMerchantCommand *CreateMerchantCommand) ToString() string {
	bytes, _ := json.Marshal(createMerchantCommand)
	return string(bytes)
}
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
func (createMerchantRequest *CreateMerchantRequest) Validate() []string {
	var err []string

	if strings.TrimSpace(createMerchantRequest.Name) == "" {
		err = append(err, "name can not be empty")
	}
	if createMerchantRequest.DiscountPercent < 0  || createMerchantRequest.DiscountPercent > 100 {
		err = append(err, "invalid discount_percent")
	}

	return err
}
func (createMerchantRequest *CreateMerchantRequest) BuildCommand() CreateMerchantCommand {
	createMerchantCommand := CreateMerchantCommand {
		Name: strings.TrimSpace(createMerchantRequest.Name),
		DiscountPercent: createMerchantRequest.DiscountPercent,
	}

	Logger.Info("CreateMerchantCommand :: ", createMerchantCommand.ToString())
	return createMerchantCommand
}
func (command *CreateMerchantCommand) ExecuteBusinessLogic() (*models.Merchant, errors.BusinessLogicError) {
	
	merchant := models.Merchant {
		Name: command.Name,
		DiscountPercent: null.Float64From(float64(command.DiscountPercent)),
	}
	businessError := errors.BusinessLogicError{}
	defer func() {
		if !businessError.IsNil() {
			Logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	merchants, err := Repositories.MerchantsRepository.GetMerchants("name = ? ", command.Name)
	if err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &merchant, businessError
	}
	if len(merchants) > 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "Merchant with same name already exists"

		return &merchant, businessError
	}

	if err := Repositories.MerchantsRepository.PutMerchant(&merchant, nil); err != nil {
		Logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &merchant, businessError
	}

	return &merchant, businessError
}