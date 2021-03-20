package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test_01", 0777)
	os.MkdirAll("testfiles/test_02", 0777)
	err := os.Remove("test_01")
	if err != nil {
		fmt.Println(err)
	}
	os.RemoveAll("testfiles")
}