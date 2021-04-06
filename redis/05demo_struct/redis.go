package redis

import (
	"context"
	"time"

	redisV8 "github.com/go-redis/redis/v8"
	"github.com/liuchengshun/imformal-form/redis/encrypt"
)

func NewClient() *redisV8.Client {
	opts := &redisV8.Options{
		Addr:    "localhost:6379",
	}
	return redisV8.NewClient(opts)
}

func SetEncode(ctx context.Context, key string, val string, expiration time.Duration) error {
	rdb := NewClient()
	valEn := encrypt.AesEncryptCFB(val)
	return rdb.Set(ctx, key, valEn, expiration).Err()
}

func GetDecode(ctx context.Context, key string) (string, error) {
	rdb := NewClient()
	valEn, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return valEn, err
	}
	return encrypt.AesDecryptCFB(valEn), nil
}
