package main

import (
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	opt, err := redis.ParseURL("redis://liuchengshun:shunshun@localhost:6379/1")
	if err != nil {
		panic(err)
	}
	fmt.Println("option:", opt)
	// rdb := redis.NewClient(opt)
}
