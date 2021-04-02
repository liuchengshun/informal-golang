package main

import (
	"context"
	"fmt"
	"strings"

	"log"
	"io"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"

	redisV8 "github.com/go-redis/redis/v8"
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

	// err := client.Set(ctx, "key", "value", 0).Err()
	// if err != nil {
	// 	panic(err)
	// }

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
	if cmd.Name() == "set" && args[1] == "key" {
		if val, ok := args[2].(string); ok {
			args[2] = aesEncryptCFB(val)
		}
	}
	return ctx, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redisV8.Cmder) error {
	args := cmd.Args()
	if cmd.Name() == "get" && args[0] == "get" && args[1] == "key" {
		fmt.Println("cmd.Args():", cmd.Args())
		fmt.Println("cmd.FullName():", cmd.FullName())
		fmt.Println("cmd.String():", cmd.String())

		vals := strings.Split(cmd.String(), ": ")
		res := aesDecryptCFB(vals[1])
		fmt.Println("res:", res)

		vals[1] = res
		str := strings.Join(vals, ": ")
		fmt.Println("str:", str)
	}
	return nil
}




var Key = []byte("0123456789ABCDEF")

func aesEncryptCFB(plainText string) (cipherStr string) {
	plainBytes := []byte(plainText)
	block, err := aes.NewCipher(Key)
	if err != nil {
		log.Fatal("create instance of encryption error:", err)
		return
	}

	cipherBytes := make([]byte, aes.BlockSize+len(plainBytes))
	iv := cipherBytes[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal("generate random iv error:", err)
		return
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherBytes[aes.BlockSize:], plainBytes)
	return hex.EncodeToString(cipherBytes)
}

func aesDecryptCFB(cipherStr string) (plainText string) {
	cipherBytes, _ := hex.DecodeString(cipherStr)
	fmt.Println("cipherBytes", cipherBytes)
	block, err := aes.NewCipher(Key)
	if err != nil {
		log.Fatal("create instance of encryption error:", err)
		return
	}

	iv := cipherBytes[:aes.BlockSize]
	plainBytes := cipherBytes[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainBytes, plainBytes)
	return string(plainBytes)
}
