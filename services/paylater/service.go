package paylater

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	. "simpl.com/loggers"

	"simpl.com/databases"
	"simpl.com/repositories"
)

type services interface {
	StartServing()
}

type simplePaylaterService struct {
	Database 			*sql.DB
	MerchantRepository  repositories.MerchantsRepository
}

func NewSimplePaylaterService() services {

	defer Logger.Info("SimplePaylaterService initialization complete")

	return simplePaylaterService{
		Database: databases.GetConnection(),
		MerchantRepository: repositories.NewMerchantsRepository(databases.GetConnection()),
	}
}

func (simplePaylaterService simplePaylaterService) StartServing() {
		
	Logger.Info("Initializing to serve SimplePaylaterService")
	router := mux.NewRouter().StrictSlash(true)

	// merchant enpoints
	router.HandleFunc("/merchant/{id}", simplePaylaterService.GetMerchantInfoEndpointHandler).Methods("GET")
	router.HandleFunc("/merchant/create", simplePaylaterService.CreateMerchantEndpointHandler).Methods("POST")
	router.HandleFunc("/merchant/update", simplePaylaterService.UpdateMerchantEndpointHandler).Methods("PATCH")
	router.HandleFunc("/merchant/report/{id}", simplePaylaterService.GenerateMerchantReportEndpointHandler).Methods("GET")

	// user endpoints
	router.HandleFunc("/user/create", simplePaylaterService.CreateUserEndpointHandler).Methods("POST")
	router.HandleFunc("/user/payback", simplePaylaterService.PaybackUserEndpointHandler).Methods("POST")
	router.HandleFunc("/user/report/{id}", simplePaylaterService.UserReportEndpointHandler).Methods("GET")

	// transaction endpoints
	router.HandleFunc("/transaction/new", simplePaylaterService.NewTransactionEndpointHandler).Methods("POST")

	Logger.Info("Serving on port 8086")
	Logger.Error(http.ListenAndServe(":8086", router))
}
