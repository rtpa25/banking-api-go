package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rtpa25/banking/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)
	if err != nil {
		WriteResponse(rw, err.Code, err.AsMessage())
	} else {
		WriteResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		WriteResponse(rw, err.Code, err.AsMessage())
	} else {
		WriteResponse(rw, http.StatusOK, customer)
	}
}

func WriteResponse(rw http.ResponseWriter, code int, data interface{}) {
	//the order should be always like this else your content type won't be applied
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(code)
	err := json.NewEncoder(rw).Encode(data)
	if err != nil {
		panic(err.Error())
	}
}
