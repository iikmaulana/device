package main

import (
	"github.com/iikmaulana/gateway/libs/helper/logger"
	"log"

	"github.com/iikmaulana/device/base-composite/config"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg := config.Init()

	serr := cfg.Start()

	err = cfg.Server.Start()
	if serr != nil {
		logger.Panic(serr)
	}
}
