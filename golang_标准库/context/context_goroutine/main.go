package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println("number:", n)
		if n == 6 {
			cancel()
			break
		}
	}
}

func gen(ctx context.Context) <-chan int {
	ch := make(chan int)
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
				time.Sleep(time.Second)
			}
		}
	}()
	return ch
}
