package service

import (
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

type DeviceRepo interface {
	AddDeviceRepo(form models.DeviceRequest) (imei string, serr serror.SError)
	UpdateDeviceRepo(imei string, form models.UpdateDeviceRequest) (serr serror.SError)
	GetDeviceByImeiRepo(imei string) (result []models.DeviceResult, serr serror.SError)
	GetAllDeviceRepo() (result models.ListDeviceResult, serr serror.SError)
	DeleteDeviceByImeiRepo(imei string) (serr serror.SError)
	AddGpsTypeRepo(form models.GpsTypeRequest) (id int64, serr serror.SError)
	UpdateGpsTypeRepo(id int64, form models.UpdateGpsTypeRequest) (serr serror.SError)
	GetGpsTypeByIDRepo(id int64) (result []models.GpsTypeResult, serr serror.SError)
	GetAllGpsTypeRepo(ndata int64, page int) (result models.ListGpsTypeResult, serr serror.SError)
	DeleteGpsTypeRepo(id int64) (serr serror.SError)
	AddHistoryRepo(form models.HistoryRequest) (id string, serr serror.SError)
	UpdateHistoryRepo(id string, form models.UpdateHistoryRequest) (serr serror.SError)
	GetHistoryByIDRepo(id string) (result []models.HistoryResult, serr serror.SError)
	GetAllHistoryRepo(ndata int64, page int) (result models.ListHistoryResult, serr serror.SError)
	DeleteHistoryRepo(id string) (serr serror.SError)
}
