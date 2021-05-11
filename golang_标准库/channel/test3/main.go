package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 5; i++ {
			done <- true
			fmt.Println("(in goroutine) number is : ", i)

			ch1 <- i

			time.Sleep(time.Second)
		}
		done <- false
	}()

	for <-done {
		data := <-ch1
		fmt.Println("data from channel:", data)
	}

	fmt.Println("End")
}
