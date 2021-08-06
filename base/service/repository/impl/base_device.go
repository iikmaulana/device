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

type deviceRepository struct {
	DB *sqlx.DB
}

func NewDeviceRepository(db *sqlx.DB) service.DeviceRepo {
	return deviceRepository{DB: db}
}

func (d deviceRepository) AddDeviceRepo(form models.DeviceRequest) (imei string, serr serror.SError) {
	err := d.DB.QueryRow(query.CreateDevice,
		form.Imei,
		form.GpstypeID,
		form.GSMNumber,
		form.Desription,
		form.OrganizationID,
		form.CreateAt,
	).Scan(&imei)

	if err != nil {
		return imei, serror.New("Can't exec query database")
	}

	return imei, nil
}

func (d deviceRepository) UpdateDeviceRepo(imei string, form models.UpdateDeviceRequest) (serr serror.SError) {
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

	prefix := ` WHERE imei = $1`
	queryVehicleID := prefix
	prefix = `UPDATE cleva.cleva_device.dv_device SET `
	queryUpdateVehicle := prefix + strings.Join(dynamicUpdate, ",") + queryVehicleID

	_, err := d.DB.Exec(queryUpdateVehicle, imei)

	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}

func (d deviceRepository) GetDeviceByImeiRepo(imei string) (result []models.DeviceResult, serr serror.SError) {
	rows, err := d.DB.Queryx(query.GetDeviceById, imei)
	if err != nil {
		return result, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		tmp := models.DeviceResult{}
		if err := rows.StructScan(&tmp); err != nil {
			return nil, serror.New("Failed scan for struct model")
		}
		result = append(result, tmp)
	}

	return result, nil
}

func (d deviceRepository) GetAllDeviceRepo(ndata int64, page int) (result []models.DeviceResult, metas models.FormMetaData, serr serror.SError) {
	rows, err := d.DB.Queryx(query.GetAllDevice)
	if err != nil {
		return result, metas, serror.New("Can't exec query database")
	}

	defer rows.Close()
	for rows.Next() {
		temp := models.DeviceResult{}
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

func (d deviceRepository) DeleteDeviceByImeiRepo(imei string) (serr serror.SError) {
	_, err := d.DB.Exec(query.DeleteDevice, imei)
	if err != nil {
		return serror.New("Can't exec query database")
	}

	return nil
}

func (d deviceRepository) CheckImei(imei string) (result bool) {
	if err := d.DB.QueryRow(query.QueryCheckImeiDevice, imei).Scan(&result); err != nil {
		return !result
	}
	return !result
}
