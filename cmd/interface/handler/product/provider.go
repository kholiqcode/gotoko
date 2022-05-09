package product

import (
	"github.com/google/wire"
	"sync"
	"toko/cmd/domain/product/repository"
	"toko/cmd/domain/product/service"
	"toko/infrastructure/database"
	"toko/pkg/auth"
)

var (
	hdl     *ProductHandlerImpl
	hdlOnce sync.Once

	svc     *service.ProductServiceImpl
	svcOnce sync.Once

	repo     *repository.ProductRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(ProductHandler), new(*ProductHandlerImpl)),
		wire.Bind(new(service.ProductService), new(*service.ProductServiceImpl)),
		wire.Bind(new(repository.ProductRepository), new(*repository.ProductRepositoryImpl)),
	)
)

func ProvideHandler(svc service.ProductService) (*ProductHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &ProductHandlerImpl{
			Svc: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo repository.ProductRepository, jwtAuth auth.JwtToken) (*service.ProductServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &service.ProductServiceImpl{
			Repo: repo,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*repository.ProductRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &repository.ProductRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
