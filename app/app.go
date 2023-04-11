package app

import (
	handlers "go-bank-api/app/handlers"
	"go-bank-api/db"
	"go-bank-api/domain"
	"go-bank-api/env"
	"go-bank-api/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()
	dbClient := db.GetDbClient(env.GetConfig())
	//wiring
	// repository := domain.NewCustomerRepositoryStub() // mock repository implementing all the methods of a normal db repository for testing purposes
	customersRepository := domain.NewCustomerRepositoryDB(dbClient)
	// accountRepository := domain.NewAccountRepositoryDB(dbClient)

	ch := handlers.CustomerHandlers{Service: service.NewCustomerService(customersRepository)}
	//define routes
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomerById).Methods(http.MethodGet)
	router.HandleFunc("/customers/create", ch.CreateCustomer).Methods(http.MethodPost)

	//starting server
	log.Fatal(http.ListenAndServe(":8080", router))
}
