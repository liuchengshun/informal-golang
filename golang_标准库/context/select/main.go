package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int)
	// go func() {
	// 	for range
	// }

	tick := time.NewTicker(time.Second)

	for {
		select {
		case <-tick.C:
			fmt.Println("wuhu qifei")
		case <-tick.C:
			fmt.Println("heiheiheihei")
		}
	}
}
