package main

import (
	"context"
	"encoding/json"
	"fmt"
	// "strings"
	"time"

	"log"
	"io"

	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"

	redisV8 "github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var rdb *redisV8.Client

type redisHook struct {
	redisV8.Hook
}

var _ redisV8.Hook = redisHook{}

type auth struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	CompanyName string `json:"company_name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Code        string `json:"code"`
}

func main() {
	rdb = redisV8.NewClient(&redisV8.Options{
		Addr:     "localhost:6379", // redis地址
	})
	ctx := context.Background()
	rdb.AddHook(redisHook{})

	err := rdb.Set(ctx, "ryze_token", "tokenToken", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "ryze_token").Result()
	// 检测，查询是否出错
	if err != nil {
		panic(err)
	}
	fmt.Println("ryze_token", val)

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
	err = rdb.Set(ctx, id, value, 24*time.Hour).Err()
	if err != nil {
		fmt.Println("rdb.Set error:", err)
		return
	}

	val, _ = rdb.Get(ctx, id).Result()
	fmt.Println("person0:", val)

	val, _ = rdb.Get(ctx, "4455ff4d-6675-4f7e-bc76-c88525c83386").Result()
	fmt.Println("val:", val)
}

func (redisHook) BeforeProcess(ctx context.Context, cmd redisV8.Cmder) (context.Context, error) {
	if cmd.Name() == "get" && cmd.FullName() == "get" {
		args := cmd.Args()
		if args[0] == "get" && args[1] == "ryze_token" {
			valueEn, _ := rdb.GetSet(ctx, "ryze_token", "").Result()
			value := aesDecryptCFB(valueEn)
			return ctx, rdb.GetSet(ctx, "ryze_token", value).Err()
		}

		if uuidKey, ok := args[1].(string); ok {
			_, err := uuid.Parse(uuidKey)
			if err == nil && args[0] == "get" {
				valueEn, _ := rdb.GetSet(ctx, uuidKey, "").Result()
				return ctx, rdb.GetSet(ctx, uuidKey, convert(valueEn, "decode")).Err()
			}
		}
	}
	return ctx, nil
}

func (redisHook) AfterProcess(ctx context.Context, cmd redisV8.Cmder) error {
	args := cmd.Args()
	if cmd.Name() == "set" && cmd.FullName() == "set" {
		if args[0] == "set" && args[1] == "ryze_token" {
			if value, ok := args[2].(string); ok {
				valueEn := aesEncryptCFB(value)
				err := rdb.GetSet(ctx, "ryze_token", valueEn).Err()
				if err != nil {
					return err
				}
			}
		}

		if uuidKey, ok := args[1].(string); ok {
			_, err := uuid.Parse(uuidKey)
			if err == nil && args[0] == "set" {
				user := &auth{}
				if v, ok := args[2].([]byte); ok {
					if err := json.Unmarshal(v, &user); err != nil {
						return err
					}
				}
				user.Password = aesEncryptCFB(user.Password)
				valueEn, _ := json.Marshal(user)
				return rdb.GetSet(ctx, uuidKey, valueEn).Err()
			}
		}
	}

	if cmd.Name() == "get" && cmd.FullName() == "get" {
		if args[0] == "get" && args[1] == "ryze_token" {
			value, _ := rdb.GetSet(ctx, "ryze_token", "").Result()
			valueEn := aesEncryptCFB(value)
			return rdb.GetSet(ctx, "ryze_token", valueEn).Err()
		}

		if uuidKey, ok := args[1].(string); ok {
			_, err := uuid.Parse(uuidKey)
			if err == nil && args[0] == "get" {
				value, _ := rdb.GetSet(ctx, uuidKey, "").Result()
				return rdb.GetSet(ctx, uuidKey, convert(value, "encode")).Err()
			}
		}
	}
	return nil
}

func convert(v string, flag string) []byte {
	user := &auth{}
	_ = json.Unmarshal([]byte(v), user)
	if flag == "encode" {
		user.Password = aesEncryptCFB(user.Password)
		userEn, _ := json.Marshal(user)
		return userEn
	}
	if flag == "decode" {
		user.Password = aesDecryptCFB(user.Password)
		userDe, _ := json.Marshal(user)
		return userDe
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
