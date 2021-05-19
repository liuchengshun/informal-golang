package main

import (
	"container/list"
	"fmt"
)

func main() {
	nums := list.New()
	nums.PushBack(1)
	nums.PushBack(2)
	nums.PushBack(3)
	for e := nums.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
