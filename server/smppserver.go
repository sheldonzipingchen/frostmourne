package server

import (
	"bufio"
	"fmt"
	"frostmourne/config"
	"frostmourne/lib/pdu"
	"frostmourne/loglib"
	"net"

	"github.com/sirupsen/logrus"
)

var serverLog = loglib.GetLog()

func Init() {
	config := config.GetConfig()

	address := config.GetString("smpp.address")
	if address == "" {
		address = "127.0.0.1"
	}

	port := config.GetString("smpp.port")
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
	for {
		reader := bufio.NewReader(conn)
		reqPDU, err := pdu.Parse(reader)
		if err != nil {
			// 解释 pdu 出错
			serverLog.WithFields(logrus.Fields{
				"error": err,
			}).Error("Parse PDU Error")
			break
		}

		serverLog.WithFields(logrus.Fields{
			"Request PDU": reqPDU,
		}).Debug("User Request PDU")

		switch reqPDU.(type) {
		case *pdu.BindRequest:
			// todo 绑定请求
			fmt.Println("hello world")
		case *pdu.CancelSM:
			// todo 取消提交请求
			fmt.Println("hello world")
		case *pdu.SubmitSM:
			// todo 提交请求
			fmt.Println("hello world")
		case *pdu.DataSM:
			// todo 前期不支持
			fmt.Println("hello world")
		case *pdu.DeliverSM:
			fmt.Println("hello world")
		case *pdu.EnquireLink:
			fmt.Println("hello world")
		case *pdu.QuerySM:
			fmt.Println("hello world")
		case *pdu.ReplaceSM:
			fmt.Println("hello world")
		case *pdu.SubmitMulti:
			fmt.Println("hello world")
		}
	}

}
