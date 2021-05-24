package main

import "fmt"

func main() {
	number := new(int)
	fmt.Println("number:", number)
	fmt.Println("number:", *number)
}
