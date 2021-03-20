package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "test.txt"
	// create file by os.Create() with fileName
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	defer file.Close()

	for i:=0; i<10; i++ {
		file.WriteString("Just a test string 我爱中国\n")
		file.Write([]byte("Just a test byte\n"))
	}
}
