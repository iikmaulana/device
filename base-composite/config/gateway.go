package config

import (
	"os"

	"github.com/iikmaulana/gateway/controller"
	"github.com/iikmaulana/gateway/libs"
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

func (cfg *Config) initGateway() serror.SError {
	var err error

	cfg.Registry, err = controller.InitRegistry(controller.RegistryConfig{
		Address:  helper.Env(libs.AppRegistryAddr, "127.0.0.1:6379"),
		Password: helper.Env(libs.AppRegistryPwd, ""),
	})
	if err != nil {
		return serror.NewFromError(err)
	}

	cfg.Server, err = controller.NewServer(controller.ServerConfig{
		Host:      helper.Env(libs.AppHost, "127.0.0.1"),
		Port:      int(helper.StringToInt(helper.Env(libs.AppPort, "31001"), 31001)),
		Key:       helper.Env(libs.AppKey, "device"),
		Name:      helper.Env(libs.AppName, "Device"),
		Namespace: helper.Env(libs.AppNamespace, libs.NamespaceDefault),
	}, cfg.Registry)
	if err != nil {
		return serror.NewFromError(err)
	}

	cfg.Gateway = cfg.Server.AsGatewayService(os.Getenv("APP_BASEPOINT"))

	return nil
}
