package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var p []byte
	n, err := reader.Read(p)
	fmt.Printf("read byte is %d, error is %v \n", n, err)
	fmt.Printf("read value is %s \n", string(p))
}
