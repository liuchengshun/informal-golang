package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func main() {
	validate := validator.New()

	var boolTest bool = false
	err := validate.Var(boolTest, "required")
	if err != nil {
		fmt.Println(err)
		// error
		// boolTest 默认值为false 但赋值为false时相当于默认值，故error
	}

	var stringTest string = ""
	err = validate.Var(stringTest, "required")
	if err != nil {
		fmt.Println(err)
		// error
		// 当string类型被赋值空串""时，相当于默认值，故error
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