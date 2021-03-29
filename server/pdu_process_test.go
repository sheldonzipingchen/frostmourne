package server

import (
	"frostmourne/config"
	"frostmourne/frostmq"
	"frostmourne/frostredis"
	"frostmourne/loglib"
)

func init() {
	config.Init("test")
	loglib.Init()
	frostredis.Init()
	frostmq.Init()
}
