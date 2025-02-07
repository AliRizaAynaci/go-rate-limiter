package redis

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
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

func RemoveOldRequests(key string, minTimestamp int64) error {
	return client.ZRemRangeByScore(context.Background(), key, "0", strconv.FormatInt(minTimestamp, 10)).Err()
}

func AddRequest(key string, timestamp int64) error {
	err := client.ZAdd(context.Background(), key, redis.Z{
		Score:  float64(timestamp),
		Member: timestamp,
	}).Err()
	if err != nil {
		return err
	}
	return client.Expire(context.Background(), key, 60*time.Second).Err()
}

func GetRequestsCount(key string) (int64, error) {
	return client.ZCount(context.Background(), key, "-inf", "+inf").Result()
}

func GetClient() *redis.Client {
	if client == nil {
		log.Fatal("Redis client is not initialized!")
	}
	return client
}
