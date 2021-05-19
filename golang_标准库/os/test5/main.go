package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("os.Getenv", os.Getenv("GOPATH")) // D:\goproject
}
