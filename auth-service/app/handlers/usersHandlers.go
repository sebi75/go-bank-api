package handlers

import (
	"banking-auth/dto"
	"banking-auth/service"
	"banking-auth/utils"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
	Service service.UserService
}

/*
The create user handler should create a new user in the database
and and sign the newly created user, returning a JWT token
*/
func (uh UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userRequest dto.NewUserRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, err.Error())
	}
	_, appErr := uh.Service.CreateUser(userRequest)

	if appErr != nil {
		utils.ResponseWriter(w, appErr.Code, appErr.AsMessage())
		return
	}

	// utils.ResponseWriter(w, http.StatusCreated, user)
	// call the service to sign the user and return a JWT token
	//todo
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{Service: service}
}