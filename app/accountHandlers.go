package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rtpa25/banking/dto"
	"github.com/rtpa25/banking/service"
)

type AccountHandler struct {
	service service.AccountService
}

func (h *AccountHandler) NewAccount(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		customerIdInInt, err1 := strconv.Atoi(customerId)
		if err1 != nil {
			WriteResponse(rw, http.StatusBadRequest, err.Error())
		}
		request.CustomerId = int64(customerIdInInt)
		account, err := h.service.NewAccount(request)
		if err != nil {
			WriteResponse(rw, err.Code, err.Message)
		} else {
			WriteResponse(rw, http.StatusCreated, account)
		}
	}
}
