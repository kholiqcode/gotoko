package health

import "github.com/labstack/echo/v4"

type HealthHandler interface {
	GetHealth(ctx echo.Context) error
}
