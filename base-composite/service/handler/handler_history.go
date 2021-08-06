package handler

import (
	"encoding/json"
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"net/http"

	response "github.com/iikmaulana/gateway/models"
	gateway "github.com/iikmaulana/gateway/service"
)

func (ox gatewayHandler) createHisory(ctx *gateway.Context) gateway.Result {
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
	form := models.HistoryRequest{}
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

	err = ox.historyUsecase.AddHistoryUsecase(form)
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

func (ox gatewayHandler) updateHisory(ctx *gateway.Context) gateway.Result {
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
	id := ctx.Parameter("id")
	var form models.UpdateHistoryRequest
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

	err = ox.historyUsecase.UpdateHistoryUsecase(id, form)
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

func (ox gatewayHandler) deleteHisoryById(ctx *gateway.Context) gateway.Result {
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

	id := ctx.Parameter("id")
	err := ox.historyUsecase.DeleteHistoryByIdUsecase(id)
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

func (ox gatewayHandler) getHisoryById(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	id := ctx.Parameter("id")
	result, err := ox.historyUsecase.GetHistoryByIDUsecase(id)
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

func (ox gatewayHandler) getAllHisory(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	result, err := ox.historyUsecase.GetAllHistoryUsecase(10, 1)
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
