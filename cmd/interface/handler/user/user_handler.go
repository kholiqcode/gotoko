package user

import (
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Get(ctx echo.Context) error
	Detail(ctx echo.Context) error
	Create(ctx echo.Context) error
	Login(ctx echo.Context) error
	Refresh(ctx echo.Context) error
}
