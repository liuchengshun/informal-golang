package main

import (
	"fmt"
)

func main() {
	fmt.Println(For())
}

func For() int {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 && j == 3 {
				return i + j
			}
			fmt.Println("i = ", i, "j = ", j)
		}
	}
	return 0
}
