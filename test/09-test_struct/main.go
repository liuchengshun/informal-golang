package main

import "fmt"

type User struct {
	name   string
	Gender string
}

func main() {
	u := User{}

	// u.name = "liuchengshun"
	// u.Gender = "boy"
	u.AddStr("liuchengshun")

	fmt.Println("user:", u)
}

func (u *User) AddStr(name string) {
	u.name = name
}
