package redis

import (
	"context"
	"testing"
)

func BenchmarkSet(b *testing.B) {
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rdb := NewClient()
		_ = rdb.Set(ctx, "key", "value", 0).Err()
	}
}

func BenchmarkSetEncode(b *testing.B) {
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = SetEncode(ctx, "key", "value", 0)
	}
}

func BenchmarkGet(b *testing.B) {
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rdb := NewClient()
		_, _ = rdb.Get(ctx, "key").Result()
	}
}

func BenchmarkGetDecode(b *testing.B) {
	ctx := context.Background()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetDecode(ctx , "key")
	}
}
