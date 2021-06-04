package main

import "fmt"

type user struct {
	name string
}

func (u user) changeName() {
	u.name = "ZhangSan"
}

func main() {
	u := user{}

	u.name = "liu"

	u.changeName()
	fmt.Println("user:", u)
}
