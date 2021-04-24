package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// file, err := os.Open("writeAt.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// writer := bufio.NewWriter(os.Stdout)
	// writer.ReadFrom(file)
	// writer.Flush()

	reader := strings.NewReader("Go语言中文网")
	reader.Seek(-6, io.SeekEnd)
	r, _, _ := reader.ReadRune()
	fmt.Printf("%c\n", r)
}
