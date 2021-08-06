package config

import (
	"fmt"
	"github.com/iikmaulana/device/base/libs/token"
	"github.com/iikmaulana/gateway/controller"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/service"
	"github.com/jmoiron/sqlx"
)

type AppConfig interface {
	Init() Config
}

type Config struct {
	Version  string
	DB       *sqlx.DB
	Registry *controller.Registry
	Server   *service.Server
	Gateway  *service.Service
	Token    *token.Initialize
}

func Init() Config {
	var cfg Config

	errx := cfg.initGateway()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitTimezone()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitToken()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitCockroachdb()
	if errx != nil {
		errx.Panic()
	}

	errx = cfg.InitService()
	if errx != nil {
		errx.Panic()
	}

	fmt.Println("Server is running ..")
	return cfg
}

func New() Config {
	return Config{
		Version: "v1.1.5",
	}
}

func (ox *Config) Start() (errx serror.SError) {
	ch := make(chan bool)
	go func() {
		err := ox.Server.Start()
		if err != nil {
			errx = serror.NewFromErrorc(err, "Cannot starting server")
		}

		ch <- false
	}()

	<-ch
	return errx
}
