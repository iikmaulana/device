package handler

import (
	"encoding/json"
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"net/http"

	response "github.com/iikmaulana/gateway/models"
	gateway "github.com/iikmaulana/gateway/service"
)

func (ox gatewayHandler) createDevice(ctx *gateway.Context) gateway.Result {

	level := ctx.AuthorizationInfo().IsOrgAdmin

	if level != 1 && level != 2 && level != 3 {
		serr := serror.New("Token access denied")
		return ctx.JSONResponse(http.StatusUnauthorized, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	form := models.DeviceRequest{}
	err := json.Unmarshal(ctx.RawBody(), &form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusBadRequest, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	err = ox.deviceUsecase.AddDeviceUsecase(form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
	})
}

func (ox gatewayHandler) updateDevice(ctx *gateway.Context) gateway.Result {
	level := ctx.AuthorizationInfo().IsOrgAdmin

	if level != 1 && level != 2 && level != 3 {
		serr := serror.New("Token access denied")
		return ctx.JSONResponse(http.StatusUnauthorized, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}
	imei := ctx.Parameter("imei")
	var form models.UpdateDeviceRequest
	err := json.Unmarshal(ctx.RawBody(), &form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusBadRequest, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	err = ox.deviceUsecase.UpdateDeviceUsecase(imei, form)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
	})
}

func (ox gatewayHandler) getDeviceById(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	imei := ctx.Parameter("imei")
	result, err := ox.deviceUsecase.GetDeviceByImeiUsecase(imei)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
		Result:     result,
	})
}

func (ox gatewayHandler) getAllDevice(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	result, err := ox.deviceUsecase.GetAllDeviceUsecase()
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
		Result:     result,
	})
}

func (ox gatewayHandler) deleteDeviceById(ctx *gateway.Context) gateway.Result {
	level := ctx.AuthorizationInfo().IsOrgAdmin

	if level != 1 && level != 2 && level != 3 {
		serr := serror.New("Token access denied")
		return ctx.JSONResponse(http.StatusUnauthorized, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}
	imei := ctx.Parameter("imei")
	err := ox.deviceUsecase.DeleteDeviceByImeiUsecase(imei)
	if err != nil {
		serr := serror.NewFromError(err)
		return ctx.JSONResponse(http.StatusInternalServerError, response.ResponseBody{
			Error:      serr.Error(),
			Appid:      "",
			Svcid:      ox.service.Key(),
			Controller: serr.File(),
			Action:     ctx.Method(),
		})
	}

	return ctx.JSONResponse(http.StatusOK, response.ResponseBody{
		Appid:      "",
		Svcid:      ox.service.Key(),
		Controller: ctx.Path(),
		Action:     ctx.Method(),
	})
}
