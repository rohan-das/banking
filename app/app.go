package app

import (
	"github.com/rohan-das/banking/domain"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohan-das/banking/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting server
	log.Fatal(http.ListenAndServe(":8000", router))
}
