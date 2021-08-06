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

type gpsTypeRepository struct {
	DB *sqlx.DB
}

func NewGpsTypeRepository(db *sqlx.DB) service.GpsTypeRepo {
	return gpsTypeRepository{DB: db}
}

func (g gpsTypeRepository) AddGpsTypeRepo(form models.GpsTypeRequest) (id int64, serr serror.SError) {
	err := g.DB.QueryRow(query.CreateGpsType,
		form.ID,
		form.Name,
		form.Description,
		form.CreateAt,
	).Scan(&id)

	if err != nil {
		return id, serror.New("Can't exec query database")
	}

	return id, nil
}

func (g gpsTypeRepository) UpdateGpsTypeRepo(id int64, form models.UpdateGpsTypeRequest) (serr serror.SError) {
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
	prefix = `UPDATE cleva.cleva_device.dv_gpstype SET `
	queryUpdateVehicle := prefix + strings.Join(dynamicUpdate, ",") + queryVehicleID

	_, err := g.DB.Exec(queryUpdateVehicle, id)

	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}

func (g gpsTypeRepository) GetGpsTypeByIDRepo(id int64) (result []models.GpsTypeResult, serr serror.SError) {

	rows, err := g.DB.Queryx(query.GetGpsTypeById, id)
	if err != nil {
		return result, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		tmp := models.GpsTypeResult{}
		if err := rows.StructScan(&tmp); err != nil {
			return nil, serror.New("Failed scan for struct model")
		}
		result = append(result, tmp)
	}

	return result, nil
}

func (g gpsTypeRepository) GetAllGpsTypeRepo(ndata int64, page int) (result []models.GpsTypeResult, metas models.FormMetaData, serr serror.SError) {

	rows, err := g.DB.Queryx(query.GetAllGpsType)
	if err != nil {
		return result, metas, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		temp := models.GpsTypeResult{}
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

func (g gpsTypeRepository) DeleteGpsTypeRepo(id int64) (serr serror.SError) {

	_, err := g.DB.Exec(query.DeleteGpsType, id)
	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}
