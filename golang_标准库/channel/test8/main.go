package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)

	go func() {

		<-timer.C

		fmt.Println("计时器已经结束")
	}()

	time.Sleep(3 * time.Second)
	stop := timer.Stop()
	fmt.Println("stop : ", stop)
	if stop {
		fmt.Println("计时器确实已经停止了")
	}

}
