package config

import (
	"github.com/iikmaulana/device/base/controller"
	"github.com/iikmaulana/device/base/service/grpc"
	"github.com/iikmaulana/device/base/service/repository/impl"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
)

func (cfg Config) InitService() serror.SError {

	deviceRepo := impl.NewDeviceRepository(cfg.DB)
	gpsTypeRepo := impl.NewGpsTypeRepository(cfg.DB)
	historyRepo := impl.NewHistoryRepository(cfg.DB)

	deviceUsecase := controller.NewDeviceUsecase(deviceRepo)
	gpsTypeUsecase := controller.NewGpsTypeUsecase(gpsTypeRepo)
	historyUsecase := controller.NewHisotyUsecase(historyRepo)

	packets.RegisterDevicesServer(cfg.Server.Instance(), grpc.ServiceHandler(deviceUsecase, gpsTypeUsecase, historyUsecase))

	return nil
}
