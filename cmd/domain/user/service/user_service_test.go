package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"toko/cmd/domain/user/dto"
	"toko/cmd/domain/user/entity"
	"toko/cmd/domain/user/repository"
)

var userRepository = &repository.UserRepositoryMock{Mock: mock.Mock{}}
var userService = UserServiceImpl{Repo: userRepository}

func TestUserService_GetUserByIdFail(t *testing.T) {

	// program mock
	userRepository.Mock.On("Find", uint(1)).Return(nil, errors.New("user not found")).Once()

	user, err := userService.GetUserById(1)
	assert.Nil(t, user)
	assert.NotNil(t, err)
}

func TestUserService_GetUserByIdSuccess(t *testing.T) {
	user := entity.User{
		ID:       2,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	userRepository.Mock.On("Find", user.ID).Return(user, nil).Once()

	result, err := userService.GetUserById(user.ID)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Name, result.Name)
}

func TestUserService_GetUsersFail(t *testing.T) {

	// program mock
	userRepository.Mock.On("FindAll").Return(nil, errors.New("users not found")).Once()

	users, err := userService.GetUsers()
	assert.Nil(t, users)
	assert.NotNil(t, err)
}

func TestUserService_GetUsersSuccess(t *testing.T) {
	user := entity.User{
		ID:       1,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	userList := entity.UserList{&user}

	// program mock
	userRepository.Mock.On("FindAll").Return(userList, nil).Once()

	users, err := userService.GetUsers()

	assert.NotNil(t, users)
	assert.Nil(t, err)
	assert.IsType(t, dto.UserListResponse{}, *users)
	assert.Equal(t, user.ID, (*users)[0].ID)
}

func TestUserService_StoreFail(t *testing.T) {
	userReq := dto.UserRequestBody{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "bismillah",
	}

	// program mock
	userRepository.Mock.On("Insert", mock.Anything).Return(nil, errors.New("can't insert to db")).Once()

	result, err := userService.Store(&userReq)

	assert.Nil(t, result)
	assert.NotNil(t, err)
}

func TestUserService_StoreSuccess(t *testing.T) {
	user := entity.User{
		ID:       1,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}

	userReq := dto.UserRequestBody{
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}

	// program mock
	userRepository.Mock.On("Insert", mock.Anything).Return(user, nil).Once()

	result, err := userService.Store(&userReq)

	assert.NotNil(t, *result)
	assert.Nil(t, err)
	assert.IsType(t, dto.UserResponse{}, *result)
	assert.Equal(t, user.ID, result.ID)
	assert.Equal(t, user.Email, result.Email)
}
