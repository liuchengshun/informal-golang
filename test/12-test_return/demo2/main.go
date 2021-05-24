package main

import "fmt"

func testReturn() string {
	var str string = "dog"
	defer func() {
		str = "dog and cat"
		fmt.Println("I love dog and cat")
	}()
	return str
}

func main() {
	res := testReturn()
	fmt.Println("result:", res)
}
