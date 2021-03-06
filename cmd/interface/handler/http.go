package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"toko/cmd/interface/handler/cart"
	"toko/cmd/interface/handler/health"
	"toko/cmd/interface/handler/order"
	"toko/cmd/interface/handler/product"
	"toko/cmd/interface/handler/user"
	"toko/internal/protocol/http/middleware/auth"
)

type HttpHandlerImpl struct {
	user    user.UserHandler
	product product.ProductHandler
	cart    cart.CartHandler
	order   order.OrderHandler
	health  health.HealthHandler
}

func NewHttpHandler(
	user user.UserHandler,
	product product.ProductHandler,
	cart cart.CartHandler,
	order order.OrderHandler,
	health health.HealthHandler,
) *HttpHandlerImpl {
	return &HttpHandlerImpl{
		user:    user,
		product: product,
		cart:    cart,
		order:   order,
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
		authGroup.POST("/register", h.user.Create)
		authGroup.POST("/refresh", h.user.Refresh, auth.JwtVerifyRefresh())
	}

	// User group
	userGroup := e.Group("user")
	{
		userGroup.GET("", h.user.Get, auth.JwtVerifyAccess())
		userGroup.GET("/:id", h.user.Detail)
	}

	productGroup := e.Group("product")
	{
		productGroup.GET("", h.product.Get)
		productGroup.GET("/category", h.product.GetCategory)
		productGroup.GET("/:slug", h.product.Detail)
		productGroup.POST("", h.product.Create)
		productGroup.POST("/category", h.product.CreateCategory)
	}

	cartGroup := e.Group("cart")
	{
		cartGroup.GET("", h.cart.Get, auth.JwtVerifyAccess())
		cartGroup.POST("", h.cart.Add, auth.JwtVerifyAccess())
		cartGroup.PUT("", h.cart.Edit, auth.JwtVerifyAccess())
		cartGroup.DELETE("/:id", h.cart.Remove, auth.JwtVerifyAccess())
	}

	orderGroup := e.Group("order")
	{
		orderGroup.GET("", h.order.Get, auth.JwtVerifyAccess())
		orderGroup.POST("", h.order.Create, auth.JwtVerifyAccess())
	}
}
