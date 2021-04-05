package main

import (
	"encoding/json"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

type redisHook struct {
	redis.Hook
}

var _ redis.Hook = redisHook{}

type auth struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Code        string `json:"code"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:  "localhost:6379",
	})
	ctx := context.Background()

	rdb.AddHook(redisHook{})
	person0 := auth{
		Username:   "liucheng",
		Password:   "123456",
		CompanyName: "hubeidaxue",
		Email:       "123@qq.com",
		Phone:       "15123234567",
		Code:        "132456",
	}
	id := uuid.NewString()
	fmt.Println("id:", id)
	value, _ := json.Marshal(person0)
	err := rdb.Set(ctx, id, value, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("rdb.Set error:", err)
		return
	}
}

func (redisHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	return ctx, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	args := cmd.Args()
	if cmd.Name() == "set" && cmd.FullName() == "set" {
		
	}
	return nil
}
