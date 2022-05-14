//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"toko/cmd/interface/handler"
	"toko/cmd/interface/handler/health"
	"toko/cmd/interface/handler/product"
	"toko/cmd/interface/handler/user"
	"toko/config"
	"toko/infrastructure/cached"
	"toko/infrastructure/database"
	"toko/internal/protocol/http"
	"toko/internal/protocol/http/router"
	"toko/pkg/auth"
)

func InitHttpProtocol(mode string) (*http.HttpImpl, error) {
	panic(wire.Build(
		config.NewConfig,
		// Wiring for jwt
		wire.NewSet(
			wire.Bind(new(auth.JwtToken), new(*auth.JwtTokenImpl)),
			auth.NewJwtToken,
		),
		// Wiring for data storage
		wire.NewSet(
			database.NewDatabaseClient,
			cached.NewRedisClient,
		),
		// Wiring for http protocol
		wire.NewSet(
			http.NewHttpProtocol,
		),
		// Wiring for http router
		wire.NewSet(
			router.NewHttpRoute,
		),
		// Wiring for http handler
		wire.NewSet(
			handler.NewHttpHandler,
		),
		user.ProviderSet,
		product.ProviderSet,
		health.ProviderSet,
	))
}
