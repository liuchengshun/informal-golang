package main

import (
	"fmt"
	"io"
	"log"

	// "os"
	"strings"
)

func main() {
	reader_A := strings.NewReader("GO语言中文网")
	res, err := ReadFrom(reader_A, 10)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("result:", string(res))

	reader_B := strings.NewReader("湖北大学，湖北的大学")
	p := make([]byte, 100)
	n, err := reader_B.ReadAt(p, 0)
	if err != nil && err != io.EOF {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p, n)
}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}
