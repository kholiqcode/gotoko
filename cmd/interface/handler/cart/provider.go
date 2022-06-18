package cart

import (
	"github.com/google/wire"
	"sync"
	_rCart "toko/cmd/domain/cart/repository"
	_sCart "toko/cmd/domain/cart/service"
	_rProduct "toko/cmd/domain/product/repository"
	"toko/infrastructure/database"
)

var (
	hdl     *CartHandlerImpl
	hdlOnce sync.Once

	svc     *_sCart.CartServiceImpl
	svcOnce sync.Once

	repo     *_rCart.CartRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(CartHandler), new(*CartHandlerImpl)),
		wire.Bind(new(_sCart.CartService), new(*_sCart.CartServiceImpl)),
		wire.Bind(new(_rCart.CartRepository), new(*_rCart.CartRepositoryImpl)),
	)
)

func ProvideHandler(svc _sCart.CartService) (*CartHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &CartHandlerImpl{
			SvcCart: svc,
		}
	})

	return hdl, nil
}

func ProvideService(rCart _rCart.CartRepository, rProduct _rProduct.ProductRepository) (*_sCart.CartServiceImpl, error) {
	svcOnce.Do(func() {
		svc = &_sCart.CartServiceImpl{
			RepoCart:    rCart,
			RepoProduct: rProduct,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rCart.CartRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rCart.CartRepositoryImpl{
			DbMysql: db.DB,
		}
	})

	return repo, nil
}
