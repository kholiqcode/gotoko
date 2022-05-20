package product

import "github.com/labstack/echo/v4"

type ProductHandler interface {
	Get(ctx echo.Context) error
	GetCategory(ctx echo.Context) error
	Detail(ctx echo.Context) error
	Create(ctx echo.Context) error
	CreateCategory(ctx echo.Context) error
}
