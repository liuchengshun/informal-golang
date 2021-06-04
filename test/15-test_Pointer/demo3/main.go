package main

import "fmt"

type user struct {
	Name string
	Age  string
}

func main() {
	oldUser := &user{
		Name: "liuchengshun",
		Age:  "12",
	}

	newUser := *oldUser
	newUser.Name = "liuluoxuan"
	fmt.Println("newUser:", newUser, "oldUser", oldUser)
}
