package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/device/base/models"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
)

func (h Handler) CreateDevice(ctx context.Context, form *packets.DevicesRequest) (output *packets.OutputDevices, err error) {
	output = &packets.OutputDevices{
		Status: 0,
	}

	request := models.DeviceRequest{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	serr := h.deviceUsecase.AddDeviceUsecase(request)
	if serr != nil {
		return output, serr
	}

	return output, nil
}

func (h Handler) UpdateDevice(ctx context.Context, form *packets.UpdateDevicesRequest) (output *packets.OutputDevices, err error) {
	output = &packets.OutputDevices{
		Status: 0,
	}

	request := models.UpdateDeviceRequest{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	serr := h.deviceUsecase.UpdateDeviceUsecase(form.DevicesID, request)
	if serr != nil {
		return output, serr
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: []byte(""),
	}

	return output, nil
}

func (h Handler) GetDevicesById(ctx context.Context, form *packets.DevicesRequestByID) (output *packets.OutputDevices, err error) {
	output = &packets.OutputDevices{
		Status: 0,
	}

	result, serr := h.deviceUsecase.GetDeviceByImeiUsecase(form.DevicesID)
	if serr != nil {
		return output, serr
	}

	byte, err := json.Marshal(result)
	if err != nil {
		return output, serror.NewFromError(err)
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: byte,
	}

	return output, nil
}

func (h Handler) GetDevicesList(ctx context.Context, form *packets.DevicesRequest) (output *packets.OutputDevices, err error) {
	output = &packets.OutputDevices{
		Status: 0,
	}

	result, serr := h.deviceUsecase.GetAllDeviceUsecase(10, 1)
	if serr != nil {
		return output, serr
	}

	result.Result.Data = result.Data

	b, err := json.Marshal(&result.Result)
	if err != nil {
		return output, serror.NewFromError(err)
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: b,
	}

	return output, nil
}

func (h Handler) DeleteDevicesByImei(ctx context.Context, form *packets.DevicesRequestByID) (output *packets.OutputDevices, err error) {
	output = &packets.OutputDevices{
		Status: 0,
	}

	serr := h.deviceUsecase.DeleteDeviceByImeiUsecase(form.DevicesID)
	if serr != nil {
		return output, serr
	}

	return output, nil
}
