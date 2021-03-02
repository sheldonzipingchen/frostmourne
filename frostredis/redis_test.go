package frostredis

import (
	"frostmourne/config"
	"testing"
)

func init() {
	config.Init("development")
}

func TestNewRedisClient(t *testing.T) {
	redisClient := NewRedisClient()

	if redisClient == nil {
		t.Error("Redis Client is nil")
	}
}
