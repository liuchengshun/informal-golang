package main

import (
	"fmt"
	"sync"
	"time"
)

var mt sync.Mutex

var wg sync.WaitGroup

var ticket int = 10

func main() {
	wg.Add(3)

	go SallTicket(1)
	go SallTicket(2)
	go SallTicket(3)

	wg.Wait()
	fmt.Println("End")
}

func SallTicket(dep int) {
	defer wg.Done()
	for {
		mt.Lock()
		if ticket > 0 {

			time.Sleep(time.Second)
			ticket--
			fmt.Printf("出票口%d,售票一张, 还剩%d张票\n", dep, ticket)

		} else {
			mt.Unlock()
			fmt.Println("已经售罄")
			break
		}
		mt.Unlock()
	}
}
