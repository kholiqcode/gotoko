package http

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"os/signal"
	"toko/config"
	"toko/internal/protocol/http/router"
	"strconv"
	"syscall"
	"time"
)

type HttpImpl struct {
	HttpRouter *router.HttpRouterImpl
	Config     *config.Config
}

func NewHttpProtocol(HttpRouter *router.HttpRouterImpl, Config *config.Config) *HttpImpl {
	return &HttpImpl{
		HttpRouter: HttpRouter,
		Config:     Config,
	}
}

func (p *HttpImpl) setupRouter(e *echo.Echo) {
	p.HttpRouter.Router(e)
}

func (p *HttpImpl) Listen() {
	// Echo instance
	e := echo.New()

	p.setupRouter(e)

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := e.Start(":" + strconv.Itoa(p.Config.AppPort)); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout
	// of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Error().Msgf("Server forced to shutdown: %v", err)
	}
	log.Info().Msg("Server exiting:")

}
