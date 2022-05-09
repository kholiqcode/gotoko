package repository

import (
	"gorm.io/gorm"
	"toko/cmd/domain/user/entity"
)

type UserRepositoryImpl struct {
	// inject db impl to RepositoriesImpl event the db is being used by the child struct impl
	Db *gorm.DB
}

func (r UserRepositoryImpl) FindAll() (*entity.UserList, error) {
	var users entity.UserList

	if e := r.Db.Debug().Find(&users).Error; e != nil {
		return nil, e
	}

	return &users, nil
}

func (r UserRepositoryImpl) Find(userId uint) (*entity.User, error) {
	var user entity.User

	if e := r.Db.Debug().First(&user, userId).Error; e != nil {
		return nil, e
	}

	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r UserRepositoryImpl) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	if e := r.Db.Debug().Where("email = ?", email).First(&user).Error; e != nil {
		return nil, e
	}

	return &entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r UserRepositoryImpl) Insert(user *entity.User) (*entity.User, error) {
	if e := r.Db.Debug().Create(&user).Error; e != nil {
		return nil, e
	}
	return user, nil
}
