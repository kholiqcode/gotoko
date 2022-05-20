package user

import (
	"github.com/google/wire"
	"sync"
	_rUser "toko/cmd/domain/user/repository"
	_sUser "toko/cmd/domain/user/service"
	"toko/infrastructure/database"
	"toko/pkg/auth"
)

var (
	hdl     *UserHandlerImpl
	hdlOnce sync.Once

	svc     *_sUser.UserServiceImpl
	svcOnce sync.Once

	repo     *_rUser.UserRepositoryImpl
	repoOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// bind each one of the interfaces
		wire.Bind(new(UserHandler), new(*UserHandlerImpl)),
		wire.Bind(new(_sUser.UserService), new(*_sUser.UserServiceImpl)),
		wire.Bind(new(_rUser.UserRepository), new(*_rUser.UserRepositoryImpl)),
	)
)

func ProvideHandler(svc _sUser.UserService) (*UserHandlerImpl, error) {
	hdlOnce.Do(func() {
		hdl = &UserHandlerImpl{
			SvcUser: svc,
		}
	})

	return hdl, nil
}

func ProvideService(repo _rUser.UserRepository, jwtAuth auth.JwtToken) (*_sUser.UserServiceImpl, error) {

	svcOnce.Do(func() {
		svc = &_sUser.UserServiceImpl{
			RepoUser: repo,
			JwtAuth:  jwtAuth,
		}
	})

	return svc, nil
}

func ProvideRepository(db *database.DatabaseImpl) (*_rUser.UserRepositoryImpl, error) {

	repoOnce.Do(func() {
		repo = &_rUser.UserRepositoryImpl{
			Db: db.DB,
		}
	})

	return repo, nil
}
