package app

import (
	"encoding/json"
	"net/http"

	"github.com/rtpa25/banking/dto"
	"github.com/rtpa25/banking/service"
)

type TransactionHandler struct {
	service service.TransactionService
}

func (h *TransactionHandler) NewTransaction(rw http.ResponseWriter, r *http.Request) {
	var request dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		WriteResponse(rw, http.StatusBadRequest, err.Error())
	} else {
		newTransaction, err := h.service.AddNewTransaction(request)
		if err != nil {
			WriteResponse(rw, err.Code, err.Message)
		} else {
			WriteResponse(rw, http.StatusCreated, newTransaction)
		}
	}
}
