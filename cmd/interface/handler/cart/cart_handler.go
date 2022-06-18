package cart

import "github.com/labstack/echo/v4"

type CartHandler interface {
	Get(ctx echo.Context) error
	Remove(ctx echo.Context) error
	Add(ctx echo.Context) error
	Edit(ctx echo.Context) error
}
