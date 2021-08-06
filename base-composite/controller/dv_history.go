package controller

import (
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/device/base-composite/service"
	"github.com/iikmaulana/gateway/libs/helper/serror"
)

type historyUsecase struct {
	deviceRepo service.DeviceRepo
}

func NewHistoryUsecase(deviceRepo service.DeviceRepo) service.HistoryUsecase {
	return historyUsecase{deviceRepo}
}

func (h historyUsecase) AddHistoryUsecase(form models.HistoryRequest) (serr serror.SError) {
	_, serr = h.deviceRepo.AddHistoryRepo(form)
	if serr != nil {
		return serr
	}

	return nil
}

func (h historyUsecase) UpdateHistoryUsecase(id string, form models.UpdateHistoryRequest) (serr serror.SError) {
	serr = h.deviceRepo.UpdateHistoryRepo(id, form)
	if serr != nil {
		return serr
	}
	return nil
}

func (h historyUsecase) GetHistoryByIDUsecase(id string) (result []models.HistoryResult, serr serror.SError) {
	result, serr = h.deviceRepo.GetHistoryByIDRepo(id)
	if serr != nil {
		return result, serr
	}
	return result, serr
}

func (h historyUsecase) GetAllHistoryUsecase(ndata int64, page int) (result models.ListHistoryResult, serr serror.SError) {
	result, serr = h.deviceRepo.GetAllHistoryRepo(ndata, page)
	if serr != nil {
		return result, serr
	}
	return result, serr
}

func (h historyUsecase) DeleteHistoryByIdUsecase(id string) (serr serror.SError) {
	serr = h.deviceRepo.DeleteHistoryRepo(id)
	if serr != nil {
		return serr
	}
	return nil
}
