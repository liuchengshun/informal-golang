package main

import (
	"fmt"
)

func Add(a, b int) int {
	return a + b
}

func GetDouble(a int) int {
	return Add(a,a)
}	

func main() {
	fmt.Println("End")
}