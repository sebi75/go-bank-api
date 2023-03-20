package handlers

import (
	"encoding/json"
	"go-bank-api/service"
	"net/http"
)

type Customer struct {
	Name string `json:"name" xml:"name"`
	City string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	Service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomers(w http.ResponseWriter, request *http.Request) {

	switch contentType := request.Header.Get("Content-Type"); contentType {
		case "application/json":
			w.Header().Set("Content-Type", "application/json")
			customers, error := ch.Service.GetAllCustomers()
			if error != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(customers)
		// case "application/xml":
		// 	w.Header().Set("Content-Type", "application/xml")
		// 	xml.NewEncoder(w).Encode(customers)
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