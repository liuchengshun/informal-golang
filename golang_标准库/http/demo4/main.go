package main

import (
	"fmt"
	"net/http"
)

type Serve struct{}

func (s Serve) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "hello world")
}

func main() {
	s := Serve{}

	_ = http.ListenAndServe(":5656", s)
}