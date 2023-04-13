package handlers

import (
	"banking-auth/dto"
	errs "banking-auth/error"
	"banking-auth/logger"
	"banking-auth/service"
	"banking-auth/utils"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	Service service.AuthService
}

/*
The create user handler should create a new user in the database
and and sign the newly created user, returning a JWT token
*/
func (uh AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var userRequest dto.RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&userRequest)
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, errs.NewValidationError("Invalid json body").AsMessage())
	}

	registerResponseDto, appErr := uh.Service.CreateUser(userRequest)
	if appErr != nil {
		logger.Error("Error while creating user: " + appErr.Message)
		utils.ResponseWriter(w, appErr.Code, appErr.AsMessage())
		return
	}

	utils.ResponseWriter(w, http.StatusCreated, registerResponseDto)
}

func (uh AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&loginRequest)
	if err != nil {
		utils.ResponseWriter(w, http.StatusBadRequest, errs.NewValidationError("Invalid json body").AsMessage())
		return
	}

	loginResponseDto, appErr := uh.Service.LoginUser(loginRequest)
	if appErr != nil {
		logger.Error("Error white logging user: " + appErr.Message)
		utils.ResponseWriter(w, appErr.Code, appErr.AsMessage())
		return
	}

	utils.ResponseWriter(w, http.StatusOK, loginResponseDto)
}

func NewUserHandler(service service.AuthService) AuthHandler {
	return AuthHandler{Service: service}
}