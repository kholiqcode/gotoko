package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"toko/cmd/interface/handler/health"
	"toko/cmd/interface/handler/product"
	"toko/cmd/interface/handler/user"
	"toko/internal/protocol/http/middleware/auth"
)

type HttpHandlerImpl struct {
	user    user.UserHandler
	product product.ProductHandler
	health  health.HealthHandler
}

func NewHttpHandler(
	user user.UserHandler,
	product product.ProductHandler,
	health health.HealthHandler,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		user:    user,
		product: product,
		health:  health,
	}
}

func (h *HttpHandlerImpl) RegisterPath(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", h.health.GetHealth)
	// Auth
	authGroup := e.Group("auth")
	{
		authGroup.POST("/login", h.user.Login)
		authGroup.POST("/refresh", h.user.Refresh, auth.JwtVerifyRefresh())
	}

	// User group
	userGroup := e.Group("user")
	{
		userGroup.GET("", h.user.Get, auth.JwtVerifyAccess())
		userGroup.POST("", h.user.Create)
	}

	productGroup := e.Group("product")
	{
		productGroup.GET("", h.product.Get)
		productGroup.POST("", h.product.Create)
	}
}
