package main

import (
	"errors"
	"fmt"
)

type user struct {
	name string
	Age  int
}

func Add(a, b int) (c int, err error) {
	num, err := GetNum()
	fmt.Println("num: ", num)
	return
}

func GetNum() (int, error) {
	return 10, errors.New("testing")
}

func GetUser() (u user, err error) {
	err = fmt.Errorf("getting role: %v", err)
	return
}

func main() {
	res, err := Add(3, 5)
	fmt.Println("result:", res)
	fmt.Println("error:", err)

	user, err := GetUser()
	fmt.Println("user:", user)
	fmt.Println("error:", err)
}
