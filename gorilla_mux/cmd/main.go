package main

import (
	"net/http"

	"github.com/gorilla/mux"

	v1 "github.com/liuchengshun/imformal-form/gorilla_mux/handler/v1"
)

func main() {
	r := mux.NewRouter()
	// v1.SetupAPI(r)
	// http.ListenAndServe(":8080", r)

	r.HandleFunc("/login", v1.LoginHandler)
	r.HandleFunc("/logout", v1.LogoutHandler)
	http.Handle("/", r)
	// http.ListenAndServe(":8080", nil)
}
