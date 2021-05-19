package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello web server")
}

func main() {
	http.HandleFunc("/hello", sayHello)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
