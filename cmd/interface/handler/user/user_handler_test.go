package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"toko/cmd/domain/user/dto"
	"toko/cmd/domain/user/entity"
	"toko/cmd/domain/user/service"
)

var userService = &service.UserServiceMock{Mock: mock.Mock{}}
var userHandler = UserHandlerImpl{Svc: userService}

func TestUserHandler_GetSuccess(t *testing.T) {
	user := entity.User{
		ID:       1,
		Name:     "Abdul Kholiq",
		Email:    "kholiqdev@icloud.com",
		Password: "rahasia",
	}
	userList := entity.UserList{&user}
	//program mock
	userService.Mock.On("GetUsers").Return(dto.CreateUserListResponse(&userList), nil).Once()

	dataResponse := struct {
		Message string `json:"message"`
		Data    struct {
			Users dto.UserListResponse `json:"users"`
		} `json:"data"`
	}{}

	e := echo.New()
	r := httptest.NewRequest("GET", "/user", nil)
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	userHandler.Get(ctx)
	bodyRes, _ := ioutil.ReadAll(w.Result().Body)
	var err = json.Unmarshal(bodyRes, &dataResponse)
	assert.NoError(t, err)

	assert.Equal(t, 200, w.Result().StatusCode)
	assert.EqualValues(t, dataResponse.Message, "Success")
	assert.EqualValues(t, dataResponse.Data.Users[0].Email, user.Email)

}

func TestUserHandler_CreateFail(t *testing.T) {
	//program mock
	userService.Mock.On("Store", mock.Anything).Return(nil, errors.New("fail on get service")).Once()

	e := echo.New()
	bodyReq := make(map[string]interface{})
	bodyReq["name"] = "Abdul Koliq"
	bodyReq["email"] = "kholiqdev@icloud.com"
	bodyReq["password"] = "password"
	bodyJson, _ := json.Marshal(bodyReq)
	r := httptest.NewRequest("POST", "/user", bytes.NewReader(bodyJson))
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	userHandler.Create(ctx)
	bodyRes, _ := ioutil.ReadAll(w.Result().Body)
	defer r.Body.Close()

	assert.Equal(t, 500, w.Result().StatusCode)
	assert.NotContains(t, string(bodyRes), "Success", "Not Contains Success Message")
}

func TestUserHandler_CreateSuccess(t *testing.T) {
	res := dto.UserResponse{
		ID:    1,
		Name:  "Abdul Koliq",
		Email: "kholiqdev@icloud.com",
	}
	//program mock
	userService.Mock.On("Store", mock.Anything).Return(res, nil).Once()

	e := echo.New()
	bodyReq := make(map[string]interface{})
	bodyReq["name"] = "Abdul Koliq"
	bodyReq["email"] = "kholiqdev@icloud.com"
	bodyReq["password"] = "password"
	bodyJson, _ := json.Marshal(bodyReq)
	r := httptest.NewRequest("POST", "/user", bytes.NewReader(bodyJson))
	r.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	userHandler.Create(ctx)
	bodyRes, _ := ioutil.ReadAll(w.Result().Body)
	defer r.Body.Close()

	assert.Equal(t, 201, w.Result().StatusCode)
	assert.Contains(t, string(bodyRes), "Success", "Message Contains")
	assert.Contains(t, string(bodyRes), res.Email, "Email Contains")
}
