package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/device/base/models"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
)

func (h Handler) CreateGpstype(ctx context.Context, form *packets.GpstypeRequest) (output *packets.OutputGpstype, err error) {
	output = &packets.OutputGpstype{
		Status: 0,
	}

	request := models.GpsTypeRequest{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	serr := h.gpsTypeUsecase.AddGpsTypeUsecase(request)
	if serr != nil {
		return output, serr
	}

	return output, nil
}

func (h Handler) UpdateGpstype(ctx context.Context, form *packets.UpdateGpstypeRequest) (output *packets.OutputGpstype, err error) {
	output = &packets.OutputGpstype{
		Status: 0,
	}

	request := models.UpdateGpsTypeRequest{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	serr := h.gpsTypeUsecase.UpdateGpsTypeUsecase(form.GpsID, request)
	if serr != nil {
		return output, serr
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: []byte(""),
	}

	return output, nil
}

func (h Handler) GetGpstypeById(ctx context.Context, form *packets.GpstypeRequestByID) (output *packets.OutputGpstype, err error) {
	output = &packets.OutputGpstype{
		Status: 0,
	}

	result, serr := h.gpsTypeUsecase.GetGpsTypeByIDUsecase(form.GpsID)
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

func (h Handler) GetGpstypeList(cctx context.Context, form *packets.GpstypeRequest) (output *packets.OutputGpstype, err error) {
	output = &packets.OutputGpstype{
		Status: 0,
	}

	result, serr := h.gpsTypeUsecase.GetAllGpsTypeUsecase(10, 1)
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

func (h Handler) DeleteGpstypeById(ctx context.Context, form *packets.GpstypeRequestByID) (output *packets.OutputGpstype, err error) {
	output = &packets.OutputGpstype{
		Status: 0,
	}

	serr := h.gpsTypeUsecase.DeleteGpsTypeIdUsecase(form.GpsID)
	if serr != nil {
		return output, serr
	}

	return output, nil
}
