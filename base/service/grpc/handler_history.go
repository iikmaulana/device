package grpc

import (
	"context"
	"encoding/json"
	"github.com/golang/protobuf/ptypes/any"
	"github.com/iikmaulana/device/base/models"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"github.com/iikmaulana/gateway/packets"
)

func (h Handler) CreateHistory(ctx context.Context, form *packets.HistoryRequest) (output *packets.OutputHistory, err error) {
	output = &packets.OutputHistory{
		Status: 0,
	}

	request := models.HistoryRequest{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	serr := h.historyUsecase.AddHistoryUsecase(request)
	if serr != nil {
		return output, serr
	}

	return output, nil
}

func (h Handler) UpdateHistory(ctx context.Context, form *packets.UpdateHistoryRequest) (output *packets.OutputHistory, err error) {
	output = &packets.OutputHistory{
		Status: 0,
	}

	request := models.UpdateHistoryRequest{}
	errx := json.Unmarshal(form.GetData().Value, &request)
	if errx != nil {
		return output, serror.NewFromError(errx)
	}

	serr := h.historyUsecase.UpdateHistoryUsecase(form.HistoryID, request)
	if serr != nil {
		return output, serr
	}

	output.Status = 1
	output.Data = &any.Any{
		Value: []byte(""),
	}

	return output, nil
}

func (h Handler) GetHistoryById(ctx context.Context, form *packets.HistoryRequestByID) (output *packets.OutputHistory, err error) {
	output = &packets.OutputHistory{
		Status: 0,
	}

	result, serr := h.historyUsecase.GetHistoryByIDUsecase(form.HistoryID)
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

func (h Handler) GetHistoryList(ctx context.Context, form *packets.HistoryRequest) (output *packets.OutputHistory, err error) {
	output = &packets.OutputHistory{
		Status: 0,
	}

	result, serr := h.historyUsecase.GetAllHistoryUsecase(10, 1)
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

func (h Handler) DeleteHistoryById(ctx context.Context, form *packets.HistoryRequestByID) (output *packets.OutputHistory, err error) {
	output = &packets.OutputHistory{
		Status: 0,
	}

	serr := h.historyUsecase.DeleteHistoryByIdUsecase(form.HistoryID)
	if serr != nil {
		return output, serr
	}

	return output, nil
}
