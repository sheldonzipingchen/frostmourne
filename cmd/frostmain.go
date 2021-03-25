package main

import (
	"flag"
	"frostmourne/config"
	"frostmourne/frostredis"
	"frostmourne/loglib"
	"frostmourne/server"
	"log"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		log.Fatal("Usage: server -e {mode}")
	}
	flag.Parse()
	config.Init(*environment)
	loglib.Init()
	frostredis.Init()
	server.Init()
}
