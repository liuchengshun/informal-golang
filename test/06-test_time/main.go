package main

import (
	"bytes"
	"fmt"
	"time"
)

func main() {
	// for {
	// 	fmt.Println("time is running")
	// 	time.Sleep(2 * time.Second)
	// }
	for {
		test2()
	}
}

func test2() {
	dotCount := bytes.Count([]byte("hello.world"), []byte("."))
	fmt.Println("t:", dotCount)
	t := time.Duration(dotCount)
	time.Sleep(t * time.Second)
	fmt.Println("test2 End")

}
