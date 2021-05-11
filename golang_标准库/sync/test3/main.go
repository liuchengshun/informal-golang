package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex

var wg *sync.WaitGroup

func main() {
	rwMutex = new(sync.RWMutex)
	wg = new(sync.WaitGroup)

	wg.Add(3)

	go WriteData(1)
	go ReadData(2)
	go WriteData(3)

	wg.Wait()

	fmt.Println("End")
}

func WriteData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始写: write start...")
	rwMutex.Lock()
	fmt.Println(i, "开始写: write...")
	time.Sleep(2 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "写结束")
}

func ReadData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始读: Read start...")
	rwMutex.RLock()
	fmt.Println(i, "开始读: Read...")
	time.Sleep(2 * time.Second)
	rwMutex.RUnlock()
	fmt.Println(i, "读结束")
}
