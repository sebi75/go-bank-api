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
	dbClient := db.GetDbClient(config) // use the client to initialize the repositories

	//initialize repository
	authRepository := domain.NewAuthRepostoryDB(dbClient);

	//initialize the handler and the service
	uh := handlers.NewUserHandler(service.NewUserService(authRepository))

	router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	router.HandleFunc("/auth/register", uh.Register).Methods(http.MethodPost)
	router.HandleFunc("/auth/login", uh.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", uh.Verify).Methods(http.MethodGet)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Println("Starting server on address: " + address + " and port: " + port)
	log.Fatal(http.ListenAndServe(address+":"+port, router))
}