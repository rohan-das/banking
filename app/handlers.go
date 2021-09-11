package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rohan-das/banking/service"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Set("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) getCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomerById(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
