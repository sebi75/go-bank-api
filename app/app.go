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
	// repository := domain.NewCustomerRepositoryStub()
	repository := domain.NewCustomerRepositoryDB()

	ch := handlers.CustomerHandlers{Service: service.NewCustomerService(repository)}
	//define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)

	//starting server
	log.Fatal(http.ListenAndServe(":8080", router))
}