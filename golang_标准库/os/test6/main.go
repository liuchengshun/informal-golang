package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println(os.Args[1:])

	fmt.Fprintln(os.Stdout, "testing")

	reader := strings.NewReader("my name is liuchengshun")
	temp := make([]byte, 20)
	n, err := reader.Read(temp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("temp:", string(temp), "n :", n)

	// file, err := os.Open("test.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// writer := bufio.NewWriter(os.Stdout)
	// writer.ReadFrom(file)
	// writer.Flush()

	reader = strings.NewReader("hello, world")
	temp_2, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("temp_2:", string(temp_2))

	fileData, err := ioutil.ReadFile("./test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("fileDate:", string(fileData))
}
