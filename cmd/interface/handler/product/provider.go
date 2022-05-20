package product

import (
	"github.com/google/wire"
	"sync"
	_rProduct "toko/cmd/domain/product/repository"
	_sProduct "toko/cmd/domain/product/service"
	"toko/infrastructure/database"
	"toko/pkg/auth"
)

var (
	hdl     *ProductHandlerImpl
	hdlOnce sync.Once

	svc     *_sProduct.ProductServiceImpl
	svcOnce sync.Once

	repo     *_rProduct.ProductRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(ProductHandler), new(*ProductHandlerImpl)),
		wire.Bind(new(_sProduct.ProductService), new(*_sProduct.ProductServiceImpl)),
		wire.Bind(new(_rProduct.ProductRepository), new(*_rProduct.ProductRepositoryImpl)),
	)
)

func ProvideHandler(svc _sProduct.ProductService) (*ProductHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &ProductHandlerImpl{
			SvcProduct: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rProduct.ProductRepository, jwtAuth auth.JwtToken) (*_sProduct.ProductServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &_sProduct.ProductServiceImpl{
			RepoProduct: repo,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rProduct.ProductRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rProduct.ProductRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
