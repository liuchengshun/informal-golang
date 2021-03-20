package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()

	var boolTest bool
	err := validate.Var(boolTest, "required")
	if err != nil {
		fmt.Println(err) // error
	}
	
	var stringTest string = ""
	err = validate.Var(stringTest, "required")
	if err != nil {
		fmt.Println(err) // error
	}

	var emailTest string = "test@136.com"
	err = validate.Var(emailTest, "email")
	if err != nil {
		fmt.Println(err) // success
	}

	var ipTest string = "127.0.0.0"
	err = validate.Var(ipTest, "ip")
	if err != nil {
		fmt.Println(err) // success
	}

	fmt.Println("end")
}