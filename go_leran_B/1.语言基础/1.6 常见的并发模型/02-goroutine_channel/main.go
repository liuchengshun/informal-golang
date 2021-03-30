package main

import (
	"fmt"
)

func main() {
	done := make(chan int)

	go func() {
		fmt.Println("hello, world")
		// sending of channel
		<- done
	}()
	
	// receiving of channel
	done <- 1
}