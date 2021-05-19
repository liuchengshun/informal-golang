package main

import (
	"fmt"
	"reflect"
)

type student struct {
	name string
}

func (s *student) YourName() string {
	return s.name
}

func main() {
	var stu1 student = student{
		name: "LaoWang",
	}

	val := reflect.ValueOf(stu1)

	kind := val.Kind()
	fmt.Println("kind:", kind)

	inter := val.Interface()
	fmt.Println("interface:", inter)
	if _, ok := inter.(struct{}); ok {
		fmt.Println("inter is struct")
	} else {
		fmt.Println("inter is not struct")
	}

	_, OK := inter.(string)
	if OK {
		fmt.Println("inter is string")
	} else {
		fmt.Println("inter is not string")
	}

	typ := val.Type()
	fmt.Println("stu1's type is : ", typ)

	fmt.Println("val.Type().Name():", val.Type().Name())
}
