package main

import "fmt"

func main() {
	var testSlice []string = []string{
		"liuchengshun",
		"liuchengyi",
		"zhangsan",
		"Lisi",
	}

	for k, v := range testSlice {
		fmt.Printf("the key is %v, the value is %v: \n", k, v)
		if v == "liuchengyi" {
			break
		}
	}
}
