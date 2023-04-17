package middleware

import (
	"go-bank-api/domain"
	"go-bank-api/errs"
	"go-bank-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)


type AuthMiddleware struct {
	repo domain.AuthRepository
}

func (m AuthMiddleware) AuthorizationHandler() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//get the current route from the request
			currentRoute := mux.CurrentRoute(r)
			currentRouteVars := mux.Vars(r)
			authHeader := r.Header.Get("Authorization")

			if authHeader != "" {
				token := getTokenFromHeader(authHeader)

				isAuthorized := m.repo.IsAuthorized(token, currentRoute.GetName(), currentRouteVars)
			
				if isAuthorized {
					next.ServeHTTP(w, r)
				} else {
					utils.ResponseWriter(w, http.StatusUnauthorized, errs.NewUnauthorizedError("Unauthorized").AsMessage())
				}
			} else {
				utils.ResponseWriter(w, http.StatusUnauthorized, errs.NewUnauthorizedError("Token is missing").AsMessage())
			}
		})
	}
}

func getTokenFromHeader(authHeader string) string {
	// the authorization header is in the form:
	// Bearer <token>
	return authHeader[7:]
}

func NewAuthMiddleware(repo domain.AuthRepository) AuthMiddleware {
	return AuthMiddleware{repo}
}