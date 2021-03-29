package main

import (
	"fmt"
)

func Hello(name string) string {
	fmt.Printf("hello %v, monkey \n", name)
	return name
}

func main() {
	str := Hello("Lisi")
	fmt.Println("string:", str)
	fmt.Println("End")
}