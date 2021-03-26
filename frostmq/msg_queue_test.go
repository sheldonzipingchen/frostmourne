package frostmq

import (
	"frostmourne/config"
	"testing"
)

func init() {
	config.Init("test")
	Init()
}

func TestGetFrostMQClient(t *testing.T) {
	client := GetFrostMQClient()
	if client == nil {
		t.Fatal("Get MQ Client Fatal")
	}
}