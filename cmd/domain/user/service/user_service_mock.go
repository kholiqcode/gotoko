package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"toko/cmd/domain/user/dto"
)

type UserServiceMock struct {
	Mock mock.Mock
}

func (u UserServiceMock) GetUsers() (*dto.UserListResponse, error) {
	arguments := u.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		users := arguments.Get(0).(dto.UserListResponse)
		return &users, nil
	}
}

func (u UserServiceMock) GetUserById(userId uint) (*dto.UserResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceMock) Store(request *dto.UserRequestBody) (*dto.UserResponse, error) {
	arguments := u.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		user := arguments.Get(0).(dto.UserResponse)
		return &user, nil
	}
}

func (u UserServiceMock) Login(request *dto.UserRequestLogin) (*dto.UserAuthResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceMock) Refresh(userId uint) (*dto.UserAuthResponse, error) {
	//TODO implement me
	panic("implement me")
}
