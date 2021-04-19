package v1

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login is running")
	fmt.Print("starting deal login logic mark")

	fmt.Fprintln(w, "login is running and this is response")
	w.WriteHeader(http.StatusOK)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("logout is running")
	fmt.Print("starting deal logout logic mark")

	fmt.Fprintln(w, "logouthandler is running and this is response")
	w.WriteHeader(http.StatusOK)
}

func getUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["user_id"]

	fmt.Printf("the id is %s", id)

	fmt.Fprintf(w, "get user handler is usered, the id is : %s\n", id)
	w.WriteHeader(http.StatusOK)
}