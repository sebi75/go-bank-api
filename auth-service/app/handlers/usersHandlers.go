package handlers

import (
	"banking-auth/service"
	"net/http"
)

type UserHandler struct {
	Service service.UserService
}

func (uh UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// implement
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{Service: service}
}