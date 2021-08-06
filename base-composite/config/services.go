package config

import (
	"github.com/iikmaulana/device/base-composite/controller"
	"github.com/iikmaulana/device/base-composite/service/handler"
	"github.com/iikmaulana/device/base-composite/service/repository/core"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

func (cfg Config) InitService() serror.SError {

	deviceRepo, serr := core.NewDeviceRepository(cfg.Registry)
	if serr != nil {
		return serr
	}

	deviceUsecase := controller.NewDeviceUsecase(deviceRepo)
	gpstypeUsecase := controller.NewGpsTypeUsecase(deviceRepo)
	historyUsecase := controller.NewHistoryUsecase(deviceRepo)

	handler.NewGatewayHandler(cfg.Gateway, deviceUsecase, gpstypeUsecase, historyUsecase)

	return nil
}
