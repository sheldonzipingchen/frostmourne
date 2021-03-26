package frostmq

import (
	"fmt"
	"frostmourne/config"
	"frostmourne/loglib"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type FrostMQClient interface {
}

type frostMQClient struct {
	Conn *amqp.Connection
}

var conn *amqp.Connection
var msgQueueLog = loglib.GetLog()

func Init() {
	c := config.GetConfig()

	host := c.GetString("amqp.host")
	if host == "" {
		host = "localhost"
	}
	port := c.GetString("amqp.port")
	if port == "" {
		port = "5672"
	}
	username := c.GetString("amqp.username")
	if username == "" {
		username = "guest"
	}
	password := c.GetString("amqp.password")
	if password == "" {
		password = "guest"
	}

	amqpURI := fmt.Sprintf("amqp://%s:%s@%s:%s/", username, password, host, port)
	var err error
	conn, err = amqp.Dial(amqpURI)
	if err != nil {
		msgQueueLog.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("AMQP Connected fail")
	}
}

func GetFrostMQClient() FrostMQClient {
	return &frostMQClient{
		Conn: conn,
	}
}