package controller

import (
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/device/base-composite/service"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

type gpsTypeUsecase struct {
	deviceRepo service.DeviceRepo
}

func NewGpsTypeUsecase(deviceRepo service.DeviceRepo) service.GpsTypeUsecase {
	return gpsTypeUsecase{deviceRepo}
}

func (g gpsTypeUsecase) AddGpsTypeUsecase(form models.GpsTypeRequest) (serr serror.SError) {
	_, serr = g.deviceRepo.AddGpsTypeRepo(form)
	if serr != nil {
		return serr
	}

	return nil
}

func (g gpsTypeUsecase) UpdateGpsTypeUsecase(id int64, form models.UpdateGpsTypeRequest) (serr serror.SError) {
	serr = g.deviceRepo.UpdateGpsTypeRepo(id, form)
	if serr != nil {
		return serr
	}
	return nil
}

func (g gpsTypeUsecase) GetGpsTypeByIDUsecase(id int64) (result []models.GpsTypeResult, serr serror.SError) {
	result, serr = g.deviceRepo.GetGpsTypeByIDRepo(id)
	if serr != nil {
		return result, serr
	}
	return result, serr
}

func (g gpsTypeUsecase) GetAllGpsTypeUsecase(ndata int64, page int) (result models.ListGpsTypeResult, serr serror.SError) {
	result, serr = g.deviceRepo.GetAllGpsTypeRepo(ndata, page)
	if serr != nil {
		return result, serr
	}
	return result, serr
}

func (g gpsTypeUsecase) DeleteGpsTypeIdUsecase(id int64) (serr serror.SError) {
	serr = g.deviceRepo.DeleteGpsTypeRepo(id)
	if serr != nil {
		return serr
	}
	return nil
}
