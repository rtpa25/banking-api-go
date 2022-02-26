package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rtpa25/banking/domain"
	"github.com/rtpa25/banking/service"
)

func Start() {
	router := mux.NewRouter()

	//wiring
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	//definig routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	//starting server
	err := http.ListenAndServe("localhost:8080", router)
	if err != nil {
		log.Fatal(err.Error())
	}
}
