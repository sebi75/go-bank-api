package app

import (
	handlers "go-bank-api/app/handlers"
	"go-bank-api/domain"
	"go-bank-api/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	//wiring
	ch := handlers.CustomerHandlers{Service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	//define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe(":8080", router))
}