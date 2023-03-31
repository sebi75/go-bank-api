package handlers

import (
	"encoding/json"
	"go-bank-api/domain"
	"go-bank-api/errs"
	"go-bank-api/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	DateOfBirth string `json:"dateOfBirth" xml:"dateOfBirth"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	Service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, request *http.Request) {
	log.Println("GetAllCustomers")
	switch contentType := request.Header.Get("Content-Type"); contentType {
		case "application/json":
			w.Header().Set("Content-Type", "application/json")
			customers, error := ch.Service.GetAllCustomers()
			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(customers)
		default:
			w.Header().Set("Content-Type", "application/json")
			customers, error := ch.Service.GetAllCustomers()
			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(customers)
	}
}

func (ch *CustomerHandlers) CreateCustomer(w http.ResponseWriter, request *http.Request) {
	var customer domain.Customer
	err := json.NewDecoder(request.Body).Decode(&customer)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		customErr := errs.NewBadRequestError("Invalid JSON body for customer")
		w.WriteHeader(customErr.Code)
		json.NewEncoder(w).Encode(customErr)
	}
	customerCreated, errr := ch.Service.CreateCustomer(customer)
	if err != nil {
		w.WriteHeader(errr.Code)
		json.NewEncoder(w).Encode(errr)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customerCreated)
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	customerId := vars["customer_id"]
	switch contentType := request.Header.Get("Content-Type"); contentType {
		case "application/json":
			w.Header().Set("Content-Type", "application/json")
			customer, err := ch.Service.GetCustomerById(customerId)
			if err != nil {
				w.WriteHeader(err.Code)
				json.NewEncoder(w).Encode(err)
				return
			}
			json.NewEncoder(w).Encode(customer)
		default:
			w.Header().Set("Content-Type", "application/json")
			customer, error := ch.Service.GetCustomerById(customerId)
			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(customer)
	}
}