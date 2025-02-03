// internal/redis/redis_test.go
package redis_test

import (
	"context"
	"testing"
	"time"

	"rate-limiter/internal/redis"
)

func TestIncrementInRedis(t *testing.T) {
	config := redis.RedisConfig{
		Addr: "localhost:6379",
		DB:   1,
	}
	client := redis.NewRedisClient(config)
	client.FlushDB(context.Background())

	key := "test-counter"
	val, err := redis.IncrementInRedis(key)
	if err != nil {
		t.Fatalf("IncrementInRedis hatası: %v", err)
	}
	if val != 1 {
		t.Fatalf("Beklenen değer 1, ancak %d geldi", val)
	}
}

func TestSetExpire(t *testing.T) {
	config := redis.RedisConfig{
		Addr: "localhost:6379",
		DB:   1,
	}
	client := redis.NewRedisClient(config)
	client.FlushDB(context.Background())

	key := "test-expire"
	client.Set(context.Background(), key, "deger", 0)
	err := redis.SetExpire(key, 1*time.Second)
	if err != nil {
		t.Fatalf("SetExpire hatası: %v", err)
	}
	time.Sleep(2 * time.Second)
	_, err = client.Get(context.Background(), key).Result()
	if err == nil {
		t.Fatalf("Anahtarın süresi dolmasına rağmen değer mevcut")
	}
}
