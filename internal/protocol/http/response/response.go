package response

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"toko/internal/protocol/http/errors"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(c echo.Context, httpCode int, message string, data interface{}) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(httpCode)
	res := Response{
		Code:    httpCode,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(c.Response()).Encode(res)
}

func Text(c echo.Context, httpCode int, message string) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextPlain)
	c.Response().WriteHeader(httpCode)
	c.Response().Write([]byte(message))
}

// TODO: implement response error
func Err(c echo.Context, httpCode int, err error) {
	_, ok := err.(*errors.RespError)
	if !ok {
		err = errors.InternalServerError(err.Error())
	}

	er, _ := err.(*errors.RespError)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(httpCode)
	res := Response{
		Code:    httpCode,
		Message: er.Message,
	}
	json.NewEncoder(c.Response()).Encode(res)
}
