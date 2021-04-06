package main

import (
	"encoding/json"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/go-redis/redis/v8"
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
		Addr:  "localhost:6379",
	})

	ctx := context.Background()

	rdb.Set(ctx, "session", 123456, 0)
	val, err := rdb.Get(ctx, "key").Result()
	switch {
	case err == redis.Nil :
		fmt.Println("key does not exist")
	case err != nil :
		fmt.Println("Get failed", err)
	case val == "" :
		fmt.Println("value is empty")
	}
	fmt.Println("val:", val)

	val2, err := rdb.Get(ctx, "session").Result()
	fmt.Println("val2:", val2, "error:", err)
	fmt.Printf("val2 type is %T\n", val2)

	// GetSet
	oldVal, err := rdb.GetSet(ctx, "session", "abcdefg").Result()
	fmt.Println("oldVal:", oldVal, "error:", err)
	val3, err := rdb.Get(ctx, "session").Result()
	fmt.Println("val2:", val3, "error:", err)

	// SetNx
	err = rdb.SetNX(ctx, "lanqiu", "111012", 0).Err()
	fmt.Println("error:", err)
	val4, err := rdb.Get(ctx, "lanqiu").Result()
	fmt.Println("val2:", val4, "error:", err)

	// MGet
	vals, err := rdb.MGet(ctx, "session", "lanqiu", "key").Result()
	fmt.Println("vals:", vals, "error:", err)

	// MSet
	err = rdb.MSet(ctx, "person1", "liuchengshun", "person2", "liuchengyi", "person3", "pengming").Err()
	fmt.Println("err:", err)
	vals, err = rdb.MGet(ctx, "person1", "person2", "person3").Result()
	fmt.Println("vals:", vals, "error:", err)

	// Incr,IncrBy 针对一个key的数值进行递增操作
	valLanqiu, err := rdb.Incr(ctx, "lanqiu").Result()
	fmt.Println("lanqiu value:", valLanqiu, "error:", err)

	valLanqiu2, err := rdb.IncrBy(ctx, "lanqiu", 2).Result()
	fmt.Println("valLanqiu2:", valLanqiu2, "error:", err)

	// Decr, DecrBy 针对一个key的数值进行递减操作
	err = rdb.MSet(ctx, "num1", 100, "num2", 200).Err()
	fmt.Println("MSet error:", err)
	valDecr, err := rdb.Decr(ctx, "num1").Result()
	fmt.Println("Decr error:", err, "valDecr :", valDecr)
	valDecrBy, err := rdb.DecrBy(ctx, "num2", 3).Result()
	fmt.Println("valDecrBy:", valDecrBy, "error DecrBy", err)

	err = rdb.SetNX(ctx, "num3", 600, 0).Err()
	fmt.Println("SetNX error:", err)
	valDecr, err = rdb.Decr(ctx, "num3").Result()
	fmt.Println("Decr error:", err, "valDecr num3 :", valDecr)

	// Del
	// err = rdb.Del(ctx, "num3").Err()
	// fmt.Println("Del error:", err)
	valGet, err := rdb.Get(ctx, "num3").Result()
	fmt.Println("valGet:", valGet, "error valGet:", err)

	person1ID := uuid.NewString()
	fmt.Println("person1ID:", person1ID)
	person1 := user{
		ID:    person1ID,
		Name:  "liu",
		Age:   18,
		Company: "hubeidaxue",
		Address: "hubei wuhan",
	}

	person1Json, err := json.Marshal(person1)
	if err != nil {
		panic(err)
	}
	err = rdb.Set(ctx, person1ID, person1Json, 0).Err()
	fmt.Println("Set error:", err)
	valPerson1, err := rdb.Get(ctx, person1ID).Result()
	fmt.Println("valPerson1:", valPerson1, "error:", err)
	fmt.Printf("type of valPerson1 from redis is:%T\n", valPerson1)

	person2 := &user{}
	err = json.Unmarshal([]byte(valPerson1), &person2)
	fmt.Println("Person2:", person2, "///error", err)
	fmt.Println("Person2.Name:", person2.Name)
}

