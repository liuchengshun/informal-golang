package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	idStr := fmt.Sprintf("%v", id)
	fmt.Printf("idStr type is %T, value is :%v\n", idStr, idStr)

	foo3()
}

func foo3() {
	values := []int{1, 2, 3, 5}
	for _, val := range values {
	  fmt.Printf("foo3 val = %d\n", val)
	}
}