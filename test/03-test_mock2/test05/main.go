package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
}

func NewObj(name string, age int) *Person {
	fmt.Printf("new obj of %v is created", name)
	return &Person{
		Name:  name,
		Age :  age,
	}
}

func NewMessage(id string) string {
	fmt.Println("NewMessage is running")
	return "thie is message of id"
}

func (p *Person) Send() {
	fmt.Printf("email is sended")
}

func main() {
	fmt.Println("End")
}