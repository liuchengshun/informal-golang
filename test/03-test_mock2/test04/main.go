package main

import (
	"fmt"

	"bou.ke/monkey"
)

func Add(a, b int) int {
	return a + b
}

func main() {
	monkey.Patch(Add, func(a, b int) int {
		fmt.Println("aaa")
		return a * b
	})
	result := Add(2, 2)
	fmt.Println("result:", result)
}
