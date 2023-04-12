package handlers

import (
	"encoding/json"
	"go-bank-api/domain"
	"go-bank-api/errs"
	"go-bank-api/service"
	"go-bank-api/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, request *http.Request) {
	log.Println("GetAllCustomers")
	customers, err := ch.Service.GetAllCustomers("")

	if err != nil {
		utils.ResponseWriter(w, err.Code, err.AsMessage())
	}

	utils.ResponseWriter(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) CreateCustomer(w http.ResponseWriter, request *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(request.Body).Decode(&customer)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, errs.NewBadRequestError("Invalid JSON body"))
	}
	customerCreated, errr := ch.Service.CreateCustomer(customer)
	if errr != nil {
		utils.ResponseWriter(w, errr.Code, errr.AsMessage())
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customerCreated)
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customer_id"]

	customer, err := ch.Service.GetCustomerById(customerId)
	
	if err != nil {
		utils.ResponseWriter(w, err.Code, err.AsMessage())
	}
	
	utils.ResponseWriter(w, http.StatusOK, customer)

}
