package frostredis

import (
	"context"
	"frostmourne/config"
	"testing"
	"time"
)

func init() {
	config.Init("test")
	Init()
}

func TestNewRedisClient(t *testing.T) {
	redisClient := NewRedisClient()

	if redisClient == nil {
		t.Error("Redis Client is nil")
	}
}

func TestSet(t *testing.T) {
	redisClient := NewRedisClient()

	ctx := context.Background()

	oneHour := 1 * time.Hour
	result, err := redisClient.Set(ctx, "test-key", 1, oneHour)
	if err != nil {
		t.Errorf("Redis Client Set Method Error: %v", err)
	} else {
		t.Logf("Redis Client Set Method Result: %v", result)
	}
}

func TestGet(t *testing.T) {
	redisClient := NewRedisClient()

	ctx := context.Background()

	result, err := redisClient.Get(ctx, "test-key")
	if err != nil {
		t.Errorf("Redis Client Get Method Error: %v", err)
	} else {
		t.Logf("Redis Client Get Method Result: %v", result)
	}
}
