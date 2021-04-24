package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// const input = "This is The Golang Standard Library.\nWelcome you!"
	// scanner := bufio.NewScanner(strings.NewReader(input))
	// scanner.Split(bufio.ScanWords)
	// count := 0
	// for scanner.Scan() {
	// 	count++
	// 	fmt.Printf("%s\n", scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	fmt.Fprintln(os.Stderr, "reading input:", err)
	// }
	// fmt.Println(count)

	// file, err := os.Create("scanner.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// file.WriteString("http://studygolang.com.\nIt is the home of gophers.\nIf you are studying golang, welcome you!")
	// // 将文件 offset 设置到文件开头
	// file.Seek(0, os.SEEK_SET)
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	file, err := os.Open("scanner.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	// s.Split(bufio.ScanLines)
	for s.Scan() {
		fmt.Printf("%s\n", s.Text())
	}
}
