package domain

import (
	"encoding/json"
	"go-bank-api/logger"
	"net/http"
	"net/url"
	"os"
)

type AuthRepository interface {
	IsAuthorized(token string, routeName string, routeVars map[string]string) bool
}

type RemoteAuthRepository struct {

}

func (r RemoteAuthRepository) IsAuthorized(token string, routeName string, routeVars map[string]string) bool {
	url := buildVerifyURL(token, routeName, routeVars)

	if response, err := http.Get(url); err != nil {
		logger.Error("Error while sending request to auth service: " + err.Error())
		return false
	} else {
		m := map[string]bool{}
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Error while decoding response from auth service: " + err.Error())
			return false
		}
		return m["isAuthorized"]
	}
}

func NewRemoteAuthRepository() AuthRepository {
	return RemoteAuthRepository{}
}

func buildVerifyURL(token string, routeName string, routeVars map[string]string) string {
	authServiceHost := os.Getenv("AUTH_SERVICE_HOST")
	authServicePort := os.Getenv("AUTH_SERVICE_PORT")
	u := url.URL{
		Host: authServiceHost + ":" + authServicePort,
		Path: "/auth/verify",
		Scheme: "http",
	}
	q := u.Query()
	q.Add("token", token)
	q.Add("route", routeName)
	for k, v := range routeVars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}