package main

import (
	"context"
	"fmt"

	redisV8 "github.com/go-redis/redis/v8"
)

type redisHook struct {
	redisV8.Hook
}

var _ redisV8.Hook = redisHook{}

func main() {
	fmt.Println("key:", key)

	client := redisV8.NewClient(&redisV8.Options{
		Addr:     "localhost:6379", // redis地址
	})
	ctx := context.Background()
	client.AddHook(redisHook{})

	err := client.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	// 检测，查询是否出错
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func (redisHook) BeforeProcess(ctx context.Context, cmd redisV8.Cmder) (context.Context, error) {
	if cmd.Name() == "set" {
		fmt.Println("beforeProcess is running")
		args := cmd.Args()
		fmt.Println("args[2]:", args[2])
		// if val, ok := args[2].(string); ok {
		// 	args[2] = AesEncryptCFB(val)
		// }
	}
	return ctx, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redisV8.Cmder) error {
	// decode
	fmt.Println("afterProcess is running")
	// fmt.Println("Name():", cmd.Name())
	// fmt.Println("FullName():", cmd.FullName())
	// fmt.Println("Args():", cmd.Args())
	// fmt.Println("String:", cmd.String())
	return nil
}
