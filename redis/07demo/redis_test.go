package redis

import (
	"context"
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	client, err := NewRedisClient()
	if err != nil || client == nil {
		t.Fatal("new redis client failed:", err)
	}

	// set value
	ctx := context.Background()
	err = client.Set(ctx, "HUBEI", "Wuhan", 0).Err()
	if err != nil {
		t.Fatal("redis client set failed:", err)
	}

	val, err := client.Get(ctx, "HUBEI").Result()
	if err != nil || val != "Wuhan" {
		t.Fatal("redis client get failed:", err)
	}

	if err = client.Del(ctx, "HUBEI").Err(); err != nil {
		t.Fatal("redis client del failed:", err)
	}

	val, err = client.Get(ctx, "HUBEI").Result()
	if err == nil || val != "" {
		t.Fatal("redis client get failed:", err)
	}
}
