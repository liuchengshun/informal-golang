package main

import "fmt"

func main() {
	supported := []string{"id", "name", "password", "email", "phone", "company_name", "status"}
	for k, v := range supported {
		fmt.Printf("v: %v\n", v)
		fmt.Printf("k: %v\n", k)
	}
}
