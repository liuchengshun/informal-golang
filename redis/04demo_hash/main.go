package main

import (
	"fmt"
	"context"

	"github.com/go-redis/redis/v8"
	// "github.com/google/uuid"
)

type user struct {
	ID       string
	Name     string
	Age      int
	Company  string
	Address  string
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	ctx := context.Background()

	err := rdb.HSet(ctx, "user_1", "Name", "sudongpo").Err()
	fmt.Println("error:", err)
	err = rdb.HSet(ctx, "user_1", "Age", 18).Err()
	fmt.Println("error:", err)
	err = rdb.HSet(ctx, "user_1", "Company", "hubeidaxue").Err()
	fmt.Println("error:", err)
	err = rdb.HSet(ctx, "user_1", "Address", "sichaungsheng").Err()
	fmt.Println("error:", err)

	valName, err := rdb.HGet(ctx, "user_1", "Name").Result()
	fmt.Println("valName:", valName, "///error of valName:", err)

	// HGetAll  返回所有的字段名和对应的value值
	objUser1, err := rdb.HGetAll(ctx, "user_1").Result()
	fmt.Println("objUser1:", objUser1, "///error of HGetAll:", err)

	userAge, err := rdb.HIncrBy(ctx, "user_1", "Age", 2).Result()
	fmt.Println("userAge HincrBy:", userAge, "///error:", err)

	// Hkeys 返回所有字段名
	keys, err := rdb.HKeys(ctx, "user_1").Result()
	fmt.Println("Hkeys:", keys, "error:", err)

	// HLen
	userLen, err := rdb.HLen(ctx, "user_1").Result()
	fmt.Println("userLen:", userLen, "///error:", err)

	// HMGet  根据key和多个字段名，查询多个hash字段值
	userVals, err := rdb.HMGet(ctx, "user_1", "Name", "Age", "Address").Result()
	fmt.Println("userVals:", userVals, "///error:", err)

	// HMSet 根据key和多个字段名和字段值，批量设置hash字段值
	data := map[string]interface{}{
		"Name": "libai",
		"Age":   18,
		"Address": "hangzhou",
		"company": "datang",
	}

	err = rdb.HMSet(ctx, "user_2", data).Err()
	fmt.Println("error of HMSet:", err)

	user2Val, err := rdb.HMGet(ctx, "user_2", "Name", "Age", "Address", "company").Result()
	fmt.Println("user2Val:", user2Val, "///error:", err)

	// HSetNX 如果field字段不存在，则设置hash字段值
	err = rdb.HSetNX(ctx, "user_3", "Age", 300).Err()
	fmt.Println("error of HSetNX:", err)
	user3Age, err := rdb.HIncrBy(ctx, "user_3", "Age", 2).Result()
	fmt.Println("user3Age:", user3Age, "///error:", err)
	args := rdb.HExists(ctx, "user_3", "Age").Args()
	fmt.Println(args)

	// HDel 根据key和字段名，删除hash字段，支持批量删除hash字段
	rdb.HDel(ctx, "user_2", "company")
	user2Company, err := rdb.HGet(ctx, "user_2", "company").Result()
	fmt.Println("user2Company:", user2Company, "///error:", err)

	// HExists 检测hash字段名是否存在。
	isExist, err := rdb.HExists(ctx, "user_2", "company").Result()
	fmt.Println("the user_2 company exist:", isExist, "///error:", err)
	isExist, err = rdb.HExists(ctx, "user_3", "Age").Result()
	fmt.Println("the user_3 Age exist:", isExist, "///error:", err)
}