package services

import (
	"net/http"

	"github.com/gorilla/mux"
	. "simpl.com/loggers"

	endpoints "simpl.com/endpoints"
)

type services interface {
	StartServing()
}

type simplePaylaterService struct {

}

func NewSimplePaylaterService() services {

	defer Logger.Info("SimplePaylaterService initialization complete")

	return simplePaylaterService{}
}

func (simplePaylaterService simplePaylaterService) StartServing() {
		
	Logger.Info("Initializing to serve SimplePaylaterService")
	router := mux.NewRouter().StrictSlash(true)

	// merchant enpoints
	router.HandleFunc("/merchant/create", endpoints.CreateMerchantEndpointHandler).Methods("POST")
	router.HandleFunc("/merchant/update", endpoints.UpdateMerchantEndpointHandler).Methods("PATCH")

	// user endpoints
	router.HandleFunc("/user/create", endpoints.CreateUserEndpointHandler).Methods("POST")

	// transaction endpoints
	router.HandleFunc("/transaction/new", endpoints.NewTransactionEndpointHandler).Methods("POST")

	Logger.Info("Serve attempt on port 8085")
	Logger.Error(http.ListenAndServe(":8085", router))
}