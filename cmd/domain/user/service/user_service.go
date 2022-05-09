package service

import (
	"toko/cmd/domain/user/dto"
)

type UserService interface {
	GetUsers() (*dto.UserListResponse, error)
	GetUserById(userId uint) (*dto.UserResponse, error)
	Store(request *dto.UserRequestBody) (*dto.UserResponse, error)
	Login(request *dto.UserRequestLogin) (*dto.UserAuthResponse, error)
	Refresh(userId uint) (*dto.UserAuthResponse, error)
}
