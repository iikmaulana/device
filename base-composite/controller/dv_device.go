package controller

import (
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/device/base-composite/service"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

type deviceUsecase struct {
	deviceRepo service.DeviceRepo
}

func NewDeviceUsecase(deviceRepo service.DeviceRepo) service.DeviceUsecase {
	return deviceUsecase{deviceRepo: deviceRepo}
}

func (d deviceUsecase) AddDeviceUsecase(form models.DeviceRequest) (serr serror.SError) {
	_, serr = d.deviceRepo.AddDeviceRepo(form)
	if serr != nil {
		return serr
	}

	return nil
}

func (d deviceUsecase) UpdateDeviceUsecase(imei string, form models.UpdateDeviceRequest) (serr serror.SError) {
	serr = d.deviceRepo.UpdateDeviceRepo(imei, form)
	if serr != nil {
		return serr
	}
	return nil
}

func (d deviceUsecase) GetDeviceByImeiUsecase(imei string) (result []models.DeviceResult, serr serror.SError) {
	result, serr = d.deviceRepo.GetDeviceByImeiRepo(imei)
	if serr != nil {
		return result, serr
	}
	return result, serr
}

func (d deviceUsecase) GetAllDeviceUsecase() (result models.ListDeviceResult, serr serror.SError) {
	result, serr = d.deviceRepo.GetAllDeviceRepo()
	if serr != nil {
		return result, serr
	}
	return result, serr
}

func (d deviceUsecase) DeleteDeviceByImeiUsecase(imei string) (serr serror.SError) {
	serr = d.deviceRepo.DeleteDeviceByImeiRepo(imei)
	if serr != nil {
		return serr
	}
	return nil
}
