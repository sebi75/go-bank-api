package handlers

import (
	"encoding/json"
	"go-bank-api/dto"
	"go-bank-api/errs"
	"go-bank-api/service"
	"go-bank-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	Service service.AccountService
}

func (ah *AccountHandler) NewAccount(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	//customer id comes from the path
	customerId := vars["customer_id"]

	//the rest of the arguments come from the body so we need to parse it
	var newAccountRequest dto.NewAccountRequest
	err := json.NewDecoder(request.Body).Decode(&newAccountRequest)
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, errs.NewBadRequestError("Invalid JSON body"))
		return
	}
	newAccountRequest.CustomerId = customerId

	newAccount, appErr := ah.Service.NewAccount(newAccountRequest)
	if appErr != nil {
		utils.ResponseWriter(w, appErr.Code, appErr.AsMessage())
		return
	}
	utils.ResponseWriter(w, http.StatusCreated, newAccount)
}