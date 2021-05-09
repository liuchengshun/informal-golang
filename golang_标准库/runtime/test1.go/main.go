package main

import (
	"fmt"
	"runtime"
)

func init() {
	//1.获取逻辑cpu的数量
	fmt.Println("逻辑CPU的核数：", runtime.NumCPU())
	//2.设置go程序执行的最大的：[1,256]
	n := runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(n)

	//获取goroot目录：
	fmt.Println("GOROOT-->", runtime.GOROOT())

	//获取操作系统
	fmt.Println("os/platform-->", runtime.GOOS) // GOOS--> darwin，mac系统
}

func main() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("goroutine。。。")
		}

	}()

	for i := 0; i < 4; i++ {
		//让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("main。。")
	}

	fmt.Println("End")
}
