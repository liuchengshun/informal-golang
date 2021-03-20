package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	field1 := "liu"
	field2 := "niu"

	validate := validator.New()

	err := validate.VarWithValue(field1, field2, "nefield")
	if err != nil {
		fmt.Println(err)  // success 
	}

	err = validate.VarWithValue(field1, field2, "eqfield")
	if err != nil {
		fmt.Println(err)   // error "liu" != "niu"
	}
}