package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("the testing numbers is : ", i)
			time.Sleep(time.Second)
		}
		ch <- 1

		fmt.Println("goroutine over")
	}()

	data := <-ch
	fmt.Println("data form channel: ", data)
	fmt.Println("main over")
}
