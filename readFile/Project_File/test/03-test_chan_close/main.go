package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	done := make(chan struct{})

	go func() {
		// 等待终端的输入
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	// os.Stdin.Read(make([]byte, 1))
	// close(done)
	loop:
		for {
			select {
			case <-done:
				fmt.Println("done send")
				break loop
			default:
				fmt.Println("default is running in select")
				time.Sleep(time.Second)
			}
		}
	fmt.Println("End")
}
