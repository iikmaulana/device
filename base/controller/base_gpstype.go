package controller

import (
	"fmt"
	"github.com/iikmaulana/device/base/models"
	"github.com/iikmaulana/device/base/service"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/libs/utils/uttime"
	"time"
)

type gpsTypeUsecase struct {
	gpsRepo service.GpsTypeRepo
}

func NewGpsTypeUsecase(gpsRepo service.GpsTypeRepo) service.GpsTypeUsecase {
	return gpsTypeUsecase{gpsRepo}
}

func (g gpsTypeUsecase) AddGpsTypeUsecase(form models.GpsTypeRequest) (serr serror.SError) {
	gpsType, err := g.gpsRepo.GetGpsTypeByIDRepo(form.ID)
	if err != nil {
		return err
	}
	if len(gpsType) > 0 {
		return serror.New(fmt.Sprintf("Duplicate primary key"))
	}

	dateTime := uttime.ToString(uttime.DefaultDateTimeFormat, time.Now())
	form.CreateAt = dateTime

	_, err = g.gpsRepo.AddGpsTypeRepo(form)
	if err != nil {
		return err
	}

	return nil
}

func (g gpsTypeUsecase) UpdateGpsTypeUsecase(id int64, form models.UpdateGpsTypeRequest) (serr serror.SError) {
	gpsType, err := g.gpsRepo.GetGpsTypeByIDRepo(id)
	if err != nil {
		return err
	}
	if len(gpsType) <= 0 {
		return serror.New(fmt.Sprintf("No data update or id you entered is wrong"))
	}

	dateTime := uttime.ToString(uttime.DefaultDateTimeFormat, time.Now())
	form.UpdateAt = dateTime

	err = g.gpsRepo.UpdateGpsTypeRepo(id, form)
	if err != nil {
		return err
	}

	return nil
}

func (g gpsTypeUsecase) GetGpsTypeByIDUsecase(id int64) (result []models.GpsTypeResult, serr serror.SError) {

	result, err := g.gpsRepo.GetGpsTypeByIDRepo(id)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (g gpsTypeUsecase) GetAllGpsTypeUsecase(ndata int64, page int) (result models.ListGpsTypeResult, serr serror.SError) {
	device, metas, serr := g.gpsRepo.GetAllGpsTypeRepo(ndata, page)
	if serr != nil {
		return result, serr
	}

	result.Result = metas
	result.Data = device

	return result, serr
}

func (g gpsTypeUsecase) DeleteGpsTypeIdUsecase(id int64) (serr serror.SError) {
	gpsType, err := g.gpsRepo.GetGpsTypeByIDRepo(id)
	if err != nil {
		return err
	}
	if len(gpsType) <= 0 {
		return serror.New(fmt.Sprintf("No data deleted or id you entered is wrong"))
	}

	err = g.gpsRepo.DeleteGpsTypeRepo(id)
	if err != nil {
		return err
	}
	return nil
}
