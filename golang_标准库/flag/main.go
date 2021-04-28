package main

import (
	"flag"
	"fmt"
)

var b bool

func main() {
	flag.BoolVar(&b, "b", false, "布尔类型的falg")

	flag.Parse()
	fmt.Println("bool falg value:", b)
}
