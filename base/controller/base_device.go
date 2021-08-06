package controller

import (
	"github.com/iikmaulana/device/base/models"
	"github.com/iikmaulana/device/base/service"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/libs/utils/uttime"
	"time"
)

type deviceUsecase struct {
	deviceRepo service.DeviceRepo
}

func NewDeviceUsecase(deviceRepo service.DeviceRepo) service.DeviceUsecase {
	return deviceUsecase{deviceRepo: deviceRepo}
}

func (d deviceUsecase) AddDeviceUsecase(form models.DeviceRequest) (serr serror.SError) {
	if !d.deviceRepo.CheckImei(form.Imei) {
		return serror.New("Duplicate primary key")
	}

	dateTime := uttime.ToString(uttime.DefaultDateTimeFormat, time.Now())
	form.CreateAt = dateTime

	_, err := d.deviceRepo.AddDeviceRepo(form)
	if err != nil {
		return err
	}
	return nil
}

func (d deviceUsecase) UpdateDeviceUsecase(imei string, form models.UpdateDeviceRequest) (serr serror.SError) {
	if d.deviceRepo.CheckImei(imei) {
		return serror.New("No data update or id you entered is wrong")
	}

	dateTime := uttime.ToString(uttime.DefaultDateTimeFormat, time.Now())
	form.UpdateAt = dateTime

	err := d.deviceRepo.UpdateDeviceRepo(imei, form)
	if err != nil {
		return err
	}
	return nil
}

func (d deviceUsecase) GetDeviceByImeiUsecase(imei string) (result []models.DeviceResult, serr serror.SError) {
	result, err := d.deviceRepo.GetDeviceByImeiRepo(imei)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (d deviceUsecase) GetAllDeviceUsecase(ndata int64, page int) (result models.ListDeviceResult, serr serror.SError) {
	device, metas, serr := d.deviceRepo.GetAllDeviceRepo(ndata, page)
	if serr != nil {
		return result, serr
	}

	result.Result = metas
	result.Data = device

	return result, serr
}

func (d deviceUsecase) DeleteDeviceByImeiUsecase(imei string) (serr serror.SError) {
	if d.deviceRepo.CheckImei(imei) {
		return serror.New("No data update or id you entered is wrong")
	}

	err := d.deviceRepo.DeleteDeviceByImeiRepo(imei)
	if err != nil {
		return err
	}
	return nil
}
