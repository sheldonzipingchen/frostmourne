package frostredis

import (
	"context"
	"fmt"
	"frostmourne/config"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error)
}

type redisClient struct {
	Rdb *redis.Client
}

func (rc redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) (string, error) {
	result, err := rc.Rdb.Set(ctx, key, value, expiration).Result()
	return result, err
}

func (rc redisClient) Get(ctx context.Context, key string) (string, error) {
	val, err := rc.Rdb.Get(ctx, key).Result()
	return val, err
}

func NewRedisClient() RedisClient {
	config := config.GetConfig()

	redisHost := config.GetString("redis.host")
	if redisHost == "" {
		redisHost = "127.0.0.1"
	}
	redisPort := config.GetString("redis.port")
	if redisPort == "" {
		redisPort = "6379"
	}
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	redisPassword := config.GetString("redis.password")
	redisDB := config.GetInt("redis.db")
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})

	rc := redisClient{
		Rdb: rdb,
	}

	return rc
}
