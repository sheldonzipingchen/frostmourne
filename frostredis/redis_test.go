package frostredis

import (
	"context"
	"frostmourne/config"
	"testing"
	"time"
)

func init() {
	config.Init("test")
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
	_, err := redisClient.Set(ctx, "test-key", 1, oneHour)
	if err != nil {
		t.Errorf("Redis Client Set Method Error: %v", err)
	}
}
