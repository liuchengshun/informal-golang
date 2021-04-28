package main

import "fmt"

func producer(nums ...int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			ch <- n
		}
	}()
	return ch
}

func deal(ch chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range ch {
			out <- n * n
		}
	}()
	return out
}

func main() {
	in := producer(1, 2, 3, 4)
	re := deal(in)

	for num := range re {
		fmt.Println("number:", num)
	}
}
