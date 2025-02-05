package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var client *redis.Client

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

func NewRedisClient(config RedisConfig) *redis.Client {
	client = redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Redis bağlantısı başarısız: %v", err)
	}
	return client
}

func IncrementInRedis(key string) (int64, error) {
	val, err := client.Incr(context.Background(), key).Result()
	if err != nil {
		return 0, err
	}
	return val, nil
}

func SetExpire(key string, expiration time.Duration) error {
	err := client.Expire(context.Background(), key, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
