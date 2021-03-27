package db

import (
	"context"
	"fmt"
	"frostmourne/config"
	"frostmourne/ent"
	"frostmourne/ent/migrate"
	"frostmourne/loglib"

	"github.com/sirupsen/logrus"

	_ "github.com/go-sql-driver/mysql"
)

var log = loglib.GetLog()

var client *ent.Client
var err error

func Init() {
	c := config.GetConfig()

	databaseType := c.GetString("db.type")
	ip := c.GetString("db.ip")
	port := c.GetString("db.port")
	username := c.GetString("db.username")
	password := c.GetString("db.password")
	database := c.GetString("db.database")

	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2fShanghai`, username, password, ip, port, database)

	client, err = ent.Open(databaseType, dsn)
	if err != nil {
		log.WithFields(logrus.Fields{
			"dsn":   dsn,
			"error": err,
		}).Fatal("failed connecting to mysql")
	}

	ctx := context.Background()

	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)

	if err != nil {
		log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("failed creating schema resources")
	}
}

func GetClient() *ent.Client {
	return client
}
