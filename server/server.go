package server

import (
	"fmt"
	"frostmourne/config"
	"frostmourne/loglib"
	"net"

	"github.com/sirupsen/logrus"
)

var serverLog = loglib.GetLog()

func Init() {
	config := config.GetConfig()

	address := config.GetString("server.address")
	if address == "" {
		address = "127.0.0.1"
	}

	port := config.GetString("server.port")
	if port == "" {
		port = "2775"
	}

	listenAddr := fmt.Sprintf("%s:%s", address, port)
	listen, err := net.Listen("tcp", listenAddr)
	if err != nil {
		serverLog.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("frostmourne server started failed, error")
	}

	var connections []net.Conn
	defer func() {
		for _, conn := range connections {
			conn.Close()
		}
	}()

	serverLog.WithFields(logrus.Fields{
		"address": address,
		"port":    port,
	}).Info("frostmourne server started successful, listen port")

	for {
		conn, err := listen.Accept()
		if err != nil {
			serverLog.WithFields(logrus.Fields{
				"err": err,
			}).Error("Accept failed")
		}
		go process(conn)

		connections = append(connections, conn)
		serverLog.WithFields(logrus.Fields{
			"total number": len(connections),
		}).Info("total number of connections")
	}

}

func process(conn net.Conn) {
}
