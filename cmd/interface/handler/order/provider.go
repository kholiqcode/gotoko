package order

import (
	"github.com/google/wire"
	"sync"
	_rCart "toko/cmd/domain/cart/repository"
	_rOrder "toko/cmd/domain/order/repository"
	_sOrder "toko/cmd/domain/order/service"
	"toko/infrastructure/database"
	"toko/pkg/auth"
)

var (
	hdl     *OrderHandlerImpl
	hdlOnce sync.Once

	svc     *_sOrder.OrderServiceImpl
	svcOnce sync.Once

	repo     *_rOrder.OrderRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(OrderHandler), new(*OrderHandlerImpl)),
		wire.Bind(new(_sOrder.OrderService), new(*_sOrder.OrderServiceImpl)),
		wire.Bind(new(_rOrder.OrderRepository), new(*_rOrder.OrderRepositoryImpl)),
	)
)

func ProvideHandler(svc _sOrder.OrderService) (*OrderHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &OrderHandlerImpl{
			SvcOrder: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rOrder.OrderRepository, repoCart _rCart.CartRepository, jwtAuth auth.JwtToken) (*_sOrder.OrderServiceImpl, error) {
	svcOnce.Do(func() {
		svc = &_sOrder.OrderServiceImpl{
			RepoOrder: repo,
			RepoCart:  repoCart,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rOrder.OrderRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rOrder.OrderRepositoryImpl{
			DbMysql: db.DB,
		}
	})

	return repo, nil
}
