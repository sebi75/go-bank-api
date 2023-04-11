package handlers

import (
	"encoding/json"
	"go-bank-api/dto"
	"go-bank-api/service"
	"net/http"

	"go-bank-api/utils"

	"github.com/gorilla/mux"
)

type DefaultTransactionHandlers struct {
	Service service.DefaultTransactionService
}

func (th DefaultTransactionHandlers) NewTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	
	var newTransactionRequest dto.NewTransactionRequest
	err := json.NewDecoder(r.Body).Decode(&newTransactionRequest)
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, err)
		return
	}
	newTransactionRequest.AccountId = accountId

	newTransaction, appErr := th.Service.NewTransaction(newTransactionRequest)
	if appErr != nil {
		utils.ResponseWriter(w, appErr.Code, appErr.AsMessage())
		return
	}
	utils.ResponseWriter(w, http.StatusCreated, newTransaction)
}