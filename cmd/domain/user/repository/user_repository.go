package repository

import (
	"toko/cmd/domain/user/entity"
)

type UserRepository interface {
	FindAll() (*entity.UserList, error)
	Find(userId uint) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Insert(user *entity.User) (*entity.User, error)
}
