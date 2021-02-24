package frostredis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
}

type redisClient struct {
	Rdb *redis.Client
}

func (rc *redisClient) Set(key string, value interface{}, expiration time.Duration) (string, error) {
	ctx := context.Background()
	result, err := rc.Rdb.Set(ctx, key, value, expiration).Result()
	return result, err
}

func (rc *redisClient) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := rc.Rdb.Get(ctx, key).Result()
	return val, err
}

func NewRedisClient() RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	rc := redisClient{
		Rdb: rdb,
	}

	return rc
}
