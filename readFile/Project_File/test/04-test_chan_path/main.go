package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	res := filepath.Join("hello", "world")
	fmt.Println("final path:", res)
}
