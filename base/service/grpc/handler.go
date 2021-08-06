package grpc

import (
	"github.com/iikmaulana/device/base/service"
)

type Handler struct {
	deviceUsecase  service.DeviceUsecase
	gpsTypeUsecase service.GpsTypeUsecase
	historyUsecase service.HistoryUsecase
}

func ServiceHandler(
	deviceUsecase service.DeviceUsecase,
	gpsTypeUsecase service.GpsTypeUsecase,
	historyUsecase service.HistoryUsecase) *Handler {
	return &Handler{
		deviceUsecase:  deviceUsecase,
		gpsTypeUsecase: gpsTypeUsecase,
		historyUsecase: historyUsecase,
	}
}
