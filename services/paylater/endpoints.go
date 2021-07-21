package paylater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"simpl.com/endpoints/merchant/create"
	"simpl.com/endpoints/merchant/update"
	. "simpl.com/loggers"
)

type event struct {
	ID          string `json:"ID"`
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type allEvents []event

var events = allEvents{
	{
		ID:          "1",
		Title:       "Introduction to Golang",
		Description: "Come join us for a chance to learn how golang works and get to eventually try it out",
	},
}

func (service simplePaylaterService) CreateMerchantEndpointHandler(w http.ResponseWriter, r *http.Request) {

	var createMerchantRequest createmerchant.CreateMerchantRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		Logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := createMerchantRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := createMerchantRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	createMerchantCommand := createmerchant.CreateMerchantCommand{}
	createMerchantCommand.BuildFromRequest(&createMerchantRequest)

	// business logic
	response, businessError := createMerchantCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (service simplePaylaterService) GetMerchantInfoEndpointHandler(w http.ResponseWriter, r *http.Request) {

	merchantID := mux.Vars(r)["id"]
	merchants, err := service.MerchantRepository.GetMerchants("id = ?", merchantID)
	if err != nil {
		Logger.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("I am a teacup!")
		return
	}
	if len(merchants) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Merchant not found!")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(merchants)
}
func (service simplePaylaterService) CreateUserEndpointHandler(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newEvent)
}

func (service simplePaylaterService) UpdateMerchantEndpointHandler(w http.ResponseWriter, r *http.Request) {
	var updateMerchantRequest updatemerchant.UpdateMerchantRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		Logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := updateMerchantRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := updateMerchantRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	updateMerchantCommand := updatemerchant.UpdateMerchantCommand{}
	updateMerchantCommand.BuildFromRequest(&updateMerchantRequest)

	// business logic
	response, businessError := updateMerchantCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (service simplePaylaterService) NewTransactionEndpointHandler(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, singleEvent := range events {
		if singleEvent.ID == eventID {
			singleEvent.Title = updatedEvent.Title
			singleEvent.Description = updatedEvent.Description
			events = append(events[:i], singleEvent)
			json.NewEncoder(w).Encode(singleEvent)
		}
	}
}
