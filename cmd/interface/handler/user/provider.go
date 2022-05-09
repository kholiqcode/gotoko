package user

import (
	"github.com/google/wire"
	"toko/cmd/domain/user/repository"
	"toko/cmd/domain/user/service"
	"toko/infrastructure/database"
	"toko/pkg/auth"
	"sync"
)

var (
	hdl     *UserHandlerImpl
	hdlOnce sync.Once

	svc     *service.UserServiceImpl
	svcOnce sync.Once

	repo     *repository.UserRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(UserHandler), new(*UserHandlerImpl)),
		wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	)
)

func ProvideHandler(svc service.UserService) (*UserHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &UserHandlerImpl{
			Svc: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo repository.UserRepository, jwtAuth auth.JwtToken) (*service.UserServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &service.UserServiceImpl{
			Repo:    repo,
			JwtAuth: jwtAuth,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*repository.UserRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &repository.UserRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
