package controller

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/iikmaulana/device/base/models"
	"github.com/iikmaulana/device/base/service"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/libs/utils/uttime"
	"time"
)

type historyUsecase struct {
	historyRepo service.HistoryRepo
}

func NewHisotyUsecase(historyRepo service.HistoryRepo) service.HistoryUsecase {
	return historyUsecase{historyRepo}
}

func (h historyUsecase) AddHistoryUsecase(form models.HistoryRequest) (serr serror.SError) {
	history, err := h.historyRepo.GetHistoryByIDRepo(form.ID)
	if err != nil {
		return err
	}
	if len(history) > 0 {
		return serror.New(fmt.Sprintf("Duplicate primary key"))
	}

	dateTime := uttime.ToString(uttime.DefaultDateTimeFormat, time.Now())
	form.ID = uuid.New().String()
	form.Time = dateTime
	form.CreateAt = dateTime

	_, err = h.historyRepo.AddHistoryRepo(form)
	if err != nil {
		return err
	}

	return nil
}

func (h historyUsecase) UpdateHistoryUsecase(id string, form models.UpdateHistoryRequest) (serr serror.SError) {
	history, err := h.historyRepo.GetHistoryByIDRepo(id)
	if err != nil {
		return err
	}
	if len(history) <= 0 {
		return serror.New(fmt.Sprintf("No data update or id you entered is wrong"))
	}

	dateTime := uttime.ToString(uttime.DefaultDateTimeFormat, time.Now())
	form.UpdateAt = dateTime

	err = h.historyRepo.UpdateHistoryRepo(id, form)
	if err != nil {
		return err
	}

	return nil
}

func (h historyUsecase) GetHistoryByIDUsecase(id string) (result []models.HistoryResult, serr serror.SError) {

	result, err := h.historyRepo.GetHistoryByIDRepo(id)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (h historyUsecase) GetAllHistoryUsecase(ndata int64, page int) (result models.ListHistoryResult, serr serror.SError) {
	history, metas, serr := h.historyRepo.GetAllHistoryRepo(ndata, page)
	if serr != nil {
		return result, serr
	}

	result.Result = metas
	result.Data = history

	return result, serr
}

func (h historyUsecase) DeleteHistoryByIdUsecase(id string) (serr serror.SError) {

	history, err := h.historyRepo.GetHistoryByIDRepo(id)
	if err != nil {
		return err
	}
	if len(history) <= 0 {
		return serror.New(fmt.Sprintf("No data deleted or id you entered is wrong"))
	}

	err = h.historyRepo.DeleteHistoryRepo(id)
	if err != nil {
		return err
	}
	return nil
}
