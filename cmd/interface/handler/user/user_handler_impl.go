package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"toko/cmd/domain/user/dto"
	"toko/cmd/domain/user/service"
	"toko/internal/protocol/http/response"
)

type UserHandlerImpl struct {
	Svc service.UserService
}

func (h UserHandlerImpl) Get(ctx echo.Context) error {
	users, err := h.Svc.GetUsers()

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"users": users,
	})
	return nil
}

func (h UserHandlerImpl) Detail(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	user, err := h.Svc.GetUserById(uint(id))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", user)
	return nil
}

func (h UserHandlerImpl) Create(ctx echo.Context) error {
	var userDto dto.UserRequestBody

	if err := ctx.Bind(&userDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	user, err := h.Svc.Store(&userDto)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", user)
	return nil
}

func (h UserHandlerImpl) Login(ctx echo.Context) error {
	var userDto dto.UserRequestLogin

	if err := ctx.Bind(&userDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	if err := ctx.Validate(userDto); err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	res, err := h.Svc.Login(&userDto)

	if err != nil {
		response.Err(ctx, http.StatusUnauthorized, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", res)
	return nil
}

func (h UserHandlerImpl) Refresh(ctx echo.Context) error {
	userId := ctx.Get("user_id").(float64)

	res, err := h.Svc.Refresh(uint(userId))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}
	response.Json(ctx, http.StatusOK, "Success", res)
	return nil
}
