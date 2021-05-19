package main

import "fmt"

func Add(a, b int) (c int, err error) {
	return
}

func main() {
	res, err := Add(3, 5)
	fmt.Println("result:", res)
	fmt.Println("error:", err)
}
