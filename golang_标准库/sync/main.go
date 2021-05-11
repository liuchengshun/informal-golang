package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)

	go func() {
		defer wg.Done()

		Print("第一个print")
	}()

	go func() {
		defer wg.Done()

		Print("第二个print")
	}()

	go func() {
		defer wg.Done()

		Print("第三个print")
	}()

	wg.Wait()

	fmt.Println("End")
}

func Print(obj interface{}) {
	fmt.Println(obj)
}
