package health

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
	"toko/internal/protocol/http/response"
)

type HealthHandlerImpl struct {
	startAt time.Time
	host    string
}

func (ctl *HealthHandlerImpl) GetHealth(ctx echo.Context) error {

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"uptime": time.Since(ctl.startAt).String(),
		"host":   ctl.host,
	})
	return nil
}
