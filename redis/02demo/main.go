package main

import (
	"context"
	// "fmt"

	redisV8 "github.com/go-redis/redis/v8"
)

type Client = redisV8.UniversalClient

type Options = redisV8.UniversalOptions

type redisHook struct {
	redisV8.Hook
}

var _ redisV8.Hook = redisHook{}

func (redisHook) AfterProcess(ctx context.Context, cmd redisV8.Cmder) error {
	// decode
	return nil
}

func (redisHook) BeforeProcess(ctx context.Context, cmd redisV8.Cmder) (context.Context, error) {
	// encode
	return ctx, nil
}

func main() {
	// rdb := redisV8.NewUniversalClient(*redisV8.UniversalOptions{

	// })
}

