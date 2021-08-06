package handler

import (
	"github.com/iikmaulana/device/base-composite/service"
	gateway "github.com/iikmaulana/gateway/service"
)

type gatewayHandler struct {
	service        *gateway.Service
	deviceUsecase  service.DeviceUsecase
	gpstypeUsecase service.GpsTypeUsecase
	historyUsecase service.HistoryUsecase
}

func NewGatewayHandler(svc *gateway.Service,
	deviceUsecase service.DeviceUsecase,
	gpstypeUsecase service.GpsTypeUsecase,
	historyUsecase service.HistoryUsecase,
) {
	h := gatewayHandler{
		service:        svc,
		deviceUsecase:  deviceUsecase,
		gpstypeUsecase: gpstypeUsecase,
		historyUsecase: historyUsecase,
	}

	h.initRoute()
}
