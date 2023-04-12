package app

import (
	"banking-auth/app/handlers"
	"banking-auth/db"
	"banking-auth/domain"
	"banking-auth/env"
	"banking-auth/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {

	//initialize mux router

	router := mux.NewRouter()
	config := env.GetConfig()
	_ = db.GetDbClient(config) // use the client to initialize the repositories

	//initialize repository
	userRepository := domain.NewUserRepositoryStub()

	//initialize the handler and the service
	uh := handlers.NewUserHandler(service.NewUserService(userRepository))

	router.HandleFunc("/auth/register", uh.CreateUser).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")
	log.Println("Starting server on address: " + address + " and port: " + port)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}