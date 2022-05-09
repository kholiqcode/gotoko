package response

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"toko/internal/protocol/http/errors"
)

type Response struct {
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Json(c echo.Context, httpCode int, message string, data interface{}) {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c.Response().WriteHeader(httpCode)
	res := Response{
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
func Err(c echo.Context, err error) {
	_, ok := err.(*errors.RespError)
	if !ok {
		err = errors.InternalServerError(err.Error())
	}

	er, _ := err.(*errors.RespError)
	c.Response().Header().Set("Content-Type", "application/json")
	c.Response().WriteHeader(er.Code)
	res := Response{
		Message: er.Message,
	}
	json.NewEncoder(c.Response()).Encode(res)
}
