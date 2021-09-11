package app

import (
	"log"
	"net/http"

	"github.com/rohan-das/banking/domain"

	"github.com/gorilla/mux"
	"github.com/rohan-das/banking/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomerById).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
