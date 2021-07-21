package paylater

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"simpl.com/endpoints/merchant/create"
	"simpl.com/endpoints/merchant/update"
	"simpl.com/endpoints/merchant/report"

	"simpl.com/endpoints/user/create"
	"simpl.com/endpoints/user/payback"
	"simpl.com/endpoints/user/report"

	"simpl.com/endpoints/transaction/create"


	. "simpl.com/loggers"
)

/**
Merchant Endpoints Begin
**/
func (service simplePaylaterService) CreateMerchantEndpointHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

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
	
	w.Header().Set("Content-Type", "application/json")

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

func (service simplePaylaterService) UpdateMerchantEndpointHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

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

	w.WriteHeader(http.StatusOK)
}

func (service simplePaylaterService) GenerateMerchantReportEndpointHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var merchantReportRequest merchantreport.MerchantReportRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		Logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := merchantReportRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := merchantReportRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	merchantReportCommand := merchantreport.MerchantReportCommand{}
	merchantReportCommand.BuildFromRequest(&merchantReportRequest)

	// business logic
	response, businessError := merchantReportCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusOK)
}
/**
Merchant Endpoints Ends
**/


/**
User Endpoints Begin
**/
func (service simplePaylaterService) CreateUserEndpointHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var createUserRequest createuser.CreateUserRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		Logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := createUserRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := createUserRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	createUserCommand := createuser.CreateUserCommand{}
	createUserCommand.BuildFromRequest(&createUserRequest)

	// business logic
	response, businessError := createUserCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (service simplePaylaterService) PaybackUserEndpointHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var paybackUserRequest userpayback.PaybackUserRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		Logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := paybackUserRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := paybackUserRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	paybackUserCommand := userpayback.PaybackUserCommand{}
	paybackUserCommand.BuildFromRequest(&paybackUserRequest)

	// business logic
	response, businessError := paybackUserCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (service simplePaylaterService) UserReportEndpointHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var userReportRequest userreport.UserReportRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		Logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := userReportRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := userReportRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	userReportCommand := userreport.UserReportCommand{}
	userReportCommand.BuildFromRequest(&userReportRequest)

	// business logic
	response, businessError := userReportCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusCreated)
}
/**
User Endpoints Ends
**/


/**
Transaction Endpoints begins
**/
func (service simplePaylaterService) NewTransactionEndpointHandler(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	var createTransactionRequest createtransaction.CreateTransactionRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		Logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := createTransactionRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := createTransactionRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	createTransactionCommand := createtransaction.CreateTransactionCommand{}
	createTransactionCommand.BuildFromRequest(&createTransactionRequest)

	// business logic
	response, businessError := createTransactionCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusCreated)
}
/**
Transaction Endpoints ends
**/