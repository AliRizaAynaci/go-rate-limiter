package redis_test

import (
	"context"
	"rate-limiter/internal/redis"
	"testing"
	"time"
)

func TestRedisOperations(t *testing.T) {
	config := redis.RedisConfig{
		Addr: "localhost:6379",
		DB:   2,
	}
	client := redis.NewRedisClient(config)
	client.FlushDB(context.Background())

	key := "test-request"
	timestamp := time.Now().Unix()

	if err := redis.AddRequest(key, timestamp); err != nil {
		t.Fatalf("AddRequest başarısız: %v", err)
	}

	count, err := redis.GetRequestsCount(key)
	if err != nil || count != 1 {
		t.Fatalf("Beklenen 1 request, ancak %d bulundu, hata: %v", count, err)
	}

	if err := redis.RemoveOldRequests(key, timestamp-10); err != nil {
		t.Fatalf("RemoveOldRequests başarısız: %v", err)
	}
}
