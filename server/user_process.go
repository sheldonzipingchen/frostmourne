package server

import (
	"frostmourne/frostmq"
	"frostmourne/frostredis"
	"frostmourne/user"
)

type SMPPServer struct {
	RClient  frostredis.RedisClient
	MQClient frostmq.MessageMQClient

	user.User
	sessionCount int
}

func (smppServer *SMPPServer) SetSessionCount(sessionCount int) {
	smppServer.sessionCount = sessionCount
}

func (smppServer *SMPPServer) GetSessionCount() int {
	return smppServer.sessionCount
}
