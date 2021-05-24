package main

import "fmt"

func main() {
	var d = []byte("key")

	fmt.Printf("type is %T, value is %v .\n", d, d)

	fmt.Println("string of d:", string(d))
}
