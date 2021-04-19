package v1

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupAPI(r *mux.Router) {
	r.HandleFunc("/login", LoginHandler)
	r.HandleFunc("/logout", LogoutHandler)

	userAPI := r.PathPrefix("/users").Subrouter()
	userAPI.HandleFunc("/{user_id}", getUserByIDHandler).Methods(http.MethodGet)
}
