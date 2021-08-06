package core

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/device/base-composite/service"
	"github.com/iikmaulana/gateway/controller"
	"github.com/iikmaulana/gateway/libs/helper"

	"github.com/iikmaulana/gateway/libs/helper/logger"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
	"time"
)

type deviceRepository struct {
	client packets.DevicesClient
}

func NewDeviceRepository(reg *controller.Registry) (res service.DeviceRepo, serr serror.SError) {
	obj := deviceRepository{}
	i := 3

	for {
		if i <= 0 {
			serr = serror.Newc("Cannot connect to core Device",
				"[repository][core][device] while dialing core Device")
			break
		}

		conn, err := reg.GetConnection("base-device")
		if err != nil {
			logger.Warn("[repository][core][device] while dial core Device")
			time.Sleep(1 * time.Second)

			i--
			continue
		}
		obj.client = packets.NewDevicesClient(conn)
		break
	}

	return obj, serr
}

func (d deviceRepository) AddDeviceRepo(form models.DeviceRequest) (imei string, serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return imei, serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := d.client.CreateDevice(context.Background(), &packets.DevicesRequest{
		Data: &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][deviceImei:%d] while grpc AddDeviceRepo: %+v", form.Imei, err)
		logger.Err(serrFmt)
		return imei, serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {
		imei = string(output.GetData().Value)
	}

	return imei, nil
}

func (d deviceRepository) UpdateDeviceRepo(imei string, form models.UpdateDeviceRequest) (serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := d.client.UpdateDevice(context.Background(), &packets.UpdateDevicesRequest{
		DevicesID: imei,
		Data:      &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][deviceimei:%d] while grpc UpdateDevice: %+v", imei, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {

	}

	return nil
}

func (d deviceRepository) GetDeviceByImeiRepo(imei string) (result []models.DeviceResult, serr serror.SError) {
	output, err := d.client.GetDevicesById(context.Background(), &packets.DevicesRequestByID{
		DevicesID: imei,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][deviceId:%d] while grpc GetDeviceByImeiRepo: %+v", imei, err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (d deviceRepository) GetAllDeviceRepo() (result models.ListDeviceResult, serr serror.SError) {
	output, err := d.client.GetDevicesList(context.Background(), &packets.DevicesRequest{})
	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device] while grpc GetAllDeviceRepo: %+v", err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (d deviceRepository) DeleteDeviceByImeiRepo(imei string) (serr serror.SError) {
	_, err := d.client.DeleteDevicesByImei(context.Background(), &packets.DevicesRequestByID{
		DevicesID: imei,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][deviceImei:%d] while grpc DeleteDeviceByImeiRepo: %+v", imei, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}

	return nil
}

func (d deviceRepository) AddGpsTypeRepo(form models.GpsTypeRequest) (id int64, serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return id, serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := d.client.CreateGpstype(context.Background(), &packets.GpstypeRequest{
		Data: &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][gpsType:%d] while grpc AddGpsTypeRepo: %+v", form.Name, err)
		logger.Err(serrFmt)
		return id, serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {
		id = helper.StringToInt(string(output.GetData().Value), 0)
	}

	return id, nil
}

func (d deviceRepository) UpdateGpsTypeRepo(id int64, form models.UpdateGpsTypeRequest) (serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := d.client.UpdateGpstype(context.Background(), &packets.UpdateGpstypeRequest{
		GpsID: id,
		Data:  &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][gpstypeID:%d] while grpc UpdateGpsTypeRepo: %+v", id, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {

	}

	return nil
}

func (d deviceRepository) GetGpsTypeByIDRepo(id int64) (result []models.GpsTypeResult, serr serror.SError) {
	output, err := d.client.GetGpstypeById(context.Background(), &packets.GpstypeRequestByID{
		GpsID: id,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][gpstypeId:%d] while grpc GetDeviceByImeiRepo: %+v", id, err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (d deviceRepository) GetAllGpsTypeRepo(ndata int64, page int) (result models.ListGpsTypeResult, serr serror.SError) {
	output, err := d.client.GetGpstypeList(context.Background(), &packets.GpstypeRequest{})
	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device] while grpc GetAllGpsTypeRepo: %+v", err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (d deviceRepository) DeleteGpsTypeRepo(id int64) (serr serror.SError) {
	_, err := d.client.DeleteGpstypeById(context.Background(), &packets.GpstypeRequestByID{
		GpsID: id,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][gpstypeID:%d] while grpc DeleteGpsTypeRepo: %+v", id, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}

	return nil
}

func (d deviceRepository) AddHistoryRepo(form models.HistoryRequest) (id string, serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return id, serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := d.client.CreateHistory(context.Background(), &packets.HistoryRequest{
		Data: &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][history:%d] while grpc AddHistoryRepo: %+v", form.Imei, err)
		logger.Err(serrFmt)
		return id, serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {
		id = string(output.GetData().Value)
	}

	return id, nil
}

func (d deviceRepository) UpdateHistoryRepo(id string, form models.UpdateHistoryRequest) (serr serror.SError) {
	byte, err := json.Marshal(form)
	if err != nil {
		return serror.NewFromError(err)
	}

	data := any.Any{
		Value: byte,
	}

	output, err := d.client.UpdateHistory(context.Background(), &packets.UpdateHistoryRequest{
		HistoryID: id,
		Data:      &data,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][historyID:%d] while grpc UpdateHistoryRepo: %+v", id, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}
	if output.GetStatus() == 1 {

	}

	return nil
}

func (d deviceRepository) GetHistoryByIDRepo(id string) (result []models.HistoryResult, serr serror.SError) {
	output, err := d.client.GetHistoryById(context.Background(), &packets.HistoryRequestByID{
		HistoryID: id,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][historyID:%d] while grpc GetHistoryByIDRepo: %+v", id, err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (d deviceRepository) GetAllHistoryRepo(ndata int64, page int) (result models.ListHistoryResult, serr serror.SError) {
	output, err := d.client.GetHistoryList(context.Background(), &packets.HistoryRequest{})
	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device] while grpc GetAllHistoryRepo: %+v", err)
		logger.Err(serrFmt)
		return result, serror.NewFromErrorc(err, serrFmt)
	}

	if output.GetStatus() == 1 {
		err := json.Unmarshal(output.GetData().Value, &result)
		if err != nil {
			return result, serror.NewFromError(err)
		}
	}

	return result, nil
}

func (d deviceRepository) DeleteHistoryRepo(id string) (serr serror.SError) {
	_, err := d.client.DeleteHistoryById(context.Background(), &packets.HistoryRequestByID{
		HistoryID: id,
	})

	if err != nil {
		serrFmt := fmt.Sprintf("[service][repository][core][device][historyID:%d] while grpc DeleteHistoryRepo: %+v", id, err)
		logger.Err(serrFmt)
		return serror.NewFromErrorc(err, serrFmt)
	}

	return nil
}
