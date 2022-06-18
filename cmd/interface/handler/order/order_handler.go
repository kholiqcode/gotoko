package order

import "github.com/labstack/echo/v4"

type OrderHandler interface {
	Get(ctx echo.Context) error
	Create(ctx echo.Context) error
}
