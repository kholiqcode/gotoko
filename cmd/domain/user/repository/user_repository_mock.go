package repository

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"toko/cmd/domain/user/entity"
)

type UserRepositoryMock struct {
	Mock mock.Mock
}

func (r *UserRepositoryMock) FindAll() (*entity.UserList, error) {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("not found arguments")
	} else {
		users := arguments.Get(0).(entity.UserList)
		return &users, nil
	}
}

func (r *UserRepositoryMock) Find(userId uint) (*entity.User, error) {
	arguments := r.Mock.Called(userId)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found userId")
	}
	user := arguments.Get(0).(entity.User)
	return &user, nil
}

func (r *UserRepositoryMock) FindByUsername(username string) (*entity.User, error) {
	arguments := r.Mock.Called(username)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		user := arguments.Get(0).(entity.User)
		return &user, nil
	}
}

func (r *UserRepositoryMock) FindByEmail(email string) (*entity.User, error) {
	arguments := r.Mock.Called(email)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		user := arguments.Get(0).(entity.User)
		return &user, nil
	}
}

func (r *UserRepositoryMock) Insert(user *entity.User) (*entity.User, error) {
	arguments := r.Mock.Called(user)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		user := arguments.Get(0).(entity.User)
		return &user, nil
	}
}
