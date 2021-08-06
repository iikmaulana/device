package handler

import (
	"encoding/json"
	"github.com/iikmaulana/device/base-composite/models"
	"github.com/iikmaulana/gateway/libs/helper"
	"github.com/iikmaulana/gateway/libs/helper/serror"
	"net/http"

	response "github.com/iikmaulana/gateway/models"
	gateway "github.com/iikmaulana/gateway/service"
)

func (ox gatewayHandler) createGpstype(ctx *gateway.Context) gateway.Result {

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

	form := models.GpsTypeRequest{}
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

	err = ox.gpstypeUsecase.AddGpsTypeUsecase(form)
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

func (ox gatewayHandler) updateGpstype(ctx *gateway.Context) gateway.Result {
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
	var form models.UpdateGpsTypeRequest
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

	err = ox.gpstypeUsecase.UpdateGpsTypeUsecase(helper.StringToInt(id, 0), form)
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

func (ox gatewayHandler) deleteGpstypeById(ctx *gateway.Context) gateway.Result {
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
	err := ox.gpstypeUsecase.DeleteGpsTypeIdUsecase(helper.StringToInt(id, 0))
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

func (ox gatewayHandler) getGpstypeById(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	id := ctx.Parameter("id")
	result, err := ox.gpstypeUsecase.GetGpsTypeByIDUsecase(helper.StringToInt(id, 0))
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

func (ox gatewayHandler) getAllGpstype(ctx *gateway.Context) gateway.Result {
	ctx.SetHeader("Access-Control-Allow-Origin", "*")
	result, err := ox.gpstypeUsecase.GetAllGpsTypeUsecase(10, 1)
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
