package main

import "fmt"

func main() {

	ch1 := make(chan int)
	// ch2 := make(chan<- int)
	// ch3 := make(<-chan int)
	chKey := make(chan bool)

	go WriteData(ch1)
	go ReadData(ch1, chKey)

	// data := <-ch1
	// fmt.Println("receive a number from channel:", data)
	<-chKey
	fmt.Println("Main End...")
}

func ReadData(ch <-chan int, chKey chan bool) {
	data := <-ch
	fmt.Println("only read ch data: ", data)
	chKey <- false
}

func WriteData(ch chan<- int) {
	ch <- 110
	fmt.Println("write 110 num into ch")
}

func WRData(ch chan int) {
	value, ok := <-ch
	if ok {
		fmt.Println("WRData read number : ", value)
	}
	if !ok {
		ch <- 1000
		fmt.Println("WRData write 1000 into channel")
	}
}
