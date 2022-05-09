package health

import (
	"github.com/google/wire"
	"os"
	"sync"
	"time"
)

var (
	hdl     *HealthHandlerImpl
	hdlOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,

		// bind each one of the interfaces
		wire.Bind(new(HealthHandler), new(*HealthHandlerImpl)),
	)
)

func ProvideHandler() (*HealthHandlerImpl, error) {
	hdlOnce.Do(func() {
		host, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		hdl = &HealthHandlerImpl{
			startAt: time.Now(),
			host:    host,
		}
	})

	return hdl, nil
}
