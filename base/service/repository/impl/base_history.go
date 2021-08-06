package impl

import (
	"fmt"
	"github.com/iikmaulana/device/base/models"
	"github.com/iikmaulana/device/base/service"
	"github.com/iikmaulana/device/base/service/helper"
	"github.com/iikmaulana/device/base/service/repository/query"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/jmoiron/sqlx"
	"reflect"
	"strings"
)

type historyRepository struct {
	DB *sqlx.DB
}

func NewHistoryRepository(db *sqlx.DB) service.HistoryRepo {
	return historyRepository{DB: db}
}

func (h historyRepository) AddHistoryRepo(form models.HistoryRequest) (id string, serr serror.SError) {
	err := h.DB.QueryRow(query.CreateHistory,
		form.ID,
		form.Imei,
		form.VehicleID,
		form.Description,
		form.Time,
		form.Status,
		form.OrganizationID,
		form.CreateAt,
	).Scan(&id)

	if err != nil {
		return id, serror.New("Can't exec query database")
	}

	return id, nil
}

func (h historyRepository) UpdateHistoryRepo(id string, form models.UpdateHistoryRequest) (serr serror.SError) {
	var dynamicUpdate []string
	n := 0
	x := reflect.ValueOf(form)
	num := x.NumField()
	if num <= 0 {
		return serr
	}
	for i := 0; i < num; i++ {
		coloumn := x.Type().Field(i).Tag.Get("db")
		t := x.Type().Field(i).Type
		exist := x.Field(i).Interface() != reflect.Zero(t).Interface()
		if exist {
			v := fmt.Sprint(x.Field(i).Interface())
			n = n + 1
			q := coloumn + ` = ` + `'` + v + `'`
			dynamicUpdate = append(dynamicUpdate, q)
		}
	}

	prefix := ` WHERE id = $1`
	queryVehicleID := prefix
	prefix = `UPDATE cleva.cleva_device.dv_history SET `
	queryUpdateVehicle := prefix + strings.Join(dynamicUpdate, ",") + queryVehicleID

	_, err := h.DB.Exec(queryUpdateVehicle, id)

	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}

func (h historyRepository) GetHistoryByIDRepo(id string) (result []models.HistoryResult, serr serror.SError) {

	rows, err := h.DB.Queryx(query.GetHistoryById, id)
	if err != nil {
		return result, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		tmp := models.HistoryResult{}
		if err := rows.StructScan(&tmp); err != nil {
			return nil, serror.New("Failed scan for struct model")
		}
		result = append(result, tmp)
	}

	return result, nil
}

func (h historyRepository) GetAllHistoryRepo(ndata int64, page int) (result []models.HistoryResult, metas models.FormMetaData, serr serror.SError) {

	rows, err := h.DB.Queryx(query.GetAllHistory)
	if err != nil {
		return result, metas, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		temp := models.HistoryResult{}
		if err := rows.StructScan(&temp); err != nil {
			return result, metas, serror.New("Failed scan for struct model")
		}
		result = append(result, temp)
	}

	var condition string

	var withPagination = true
	if page == 0 {
		withPagination = false
	}

	var offset int
	var paginate models.FormMetaData

	if withPagination == true {
		paginate, offset = helper.Paginate(ndata, page, 1)
		condition = condition + fmt.Sprintf(" group by lrr.loan_request_restructure_id, l.payment_type, lrsv.loan_id order by lrr.loan_request_restructure_id desc limit 1 offset %v ", offset)
	} else {
		condition = condition + " group by lrr.loan_request_restructure_id, l.payment_type, lrsv.loan_id order by lrr.loan_request_restructure_id desc "
	}

	result = result
	metas = paginate
	return result, metas, nil
}

func (h historyRepository) DeleteHistoryRepo(id string) (serr serror.SError) {

	_, err := h.DB.Exec(query.DeleteHistory, id)
	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}
