package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	redisV8 "github.com/go-redis/redis/v8"

	"gitlab.ustack.com/ued/ryze/pkg/model/cryptic"
)

type redisHook struct {
	redisV8.Hook
}

var _ redisV8.Hook = redisHook{}

func main() {
	client := redisV8.NewClient(&redisV8.Options{
		Addr:     "localhost:6379", // redis地址
	})
	ctx := context.Background()
	client.AddHook(redisHook{})

	err := client.Set(ctx, "ryze_token", "xunxunmimi", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "key").Result()
	// 检测，查询是否出错
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	// res := aesDecryptCFB(val)
	// fmt.Println("res:", res)
}

func (redisHook) BeforeProcess(ctx context.Context, cmd redisV8.Cmder) (context.Context, error) {
	args := cmd.Args()
	if cmd.Name() == "set"  {
		if args[1] == "ryze_token" {
			if val, ok := args[2].(string); ok {
				args[2] = cryptic.AesEncryptCFB(val)
			}
			fmt.Println("beforeProcess is running")
		}

		if key, ok := args[1].(string); ok {
			if _, err := uuid.Parse(key); err == nil {
				payload := struct {
					Username    string `json:"username"`
					Password    string `json:"password"`
					CompanyName string `json:"company_name"`
					Email       string `json:"email"`
					Phone       string `json:"phone"`
					Code        string `json:"code"`
				}{}
				json.Unmarshal([]byte(args[2]), &payload)
			}
		}
	}
	return ctx, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redisV8.Cmder) error {
	// args := cmd.Args()
	// if cmd.Name() == "get" && args[0] == "get" && args[1] == "key" {
	// 	fmt.Println("cmd.Args():", cmd.Args())
	// 	fmt.Println("cmd.FullName():", cmd.FullName())
	// 	fmt.Println("cmd.String():", cmd.String())

	// 	vals := strings.Split(cmd.String(), ": ")
	// 	res := cryptic.AesDecryptCFB(vals[1])
	// 	fmt.Println("res:", res)

	// 	vals[1] = res
	// 	str := strings.Join(vals, ": ")
	// 	fmt.Println("str:", str)
	// }
	return nil
}