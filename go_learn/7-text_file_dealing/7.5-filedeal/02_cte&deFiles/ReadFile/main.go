package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "../../../7.3-regexp2/main.go"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(fileName, err)
		return
	}
	defer file.Close()

	// 读取数据
	buf := make([]byte, 1024)
	for {
		n, _ := file.Read(buf)
		if n == 0 {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}