package cart

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"toko/cmd/domain/cart/dto"
	"toko/cmd/domain/cart/service"
	"toko/internal/protocol/http/response"
	"toko/pkg/database"
)

type CartHandlerImpl struct {
	SvcCart service.CartService
}

func (h CartHandlerImpl) Get(ctx echo.Context) error {
	pagination := database.NewPagination(ctx)
	userId := ctx.Get("user_id").(float64)

	carts, err := h.SvcCart.GetCarts(pagination, uint(userId))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"carts": map[string]interface{}{
			"data":       carts,
			"sort":       pagination.GetSort(),
			"page":       pagination.GetPage(),
			"page_size":  pagination.GetLimit(),
			"total_page": pagination.GetTotalPage(),
			"total_rows": pagination.GetTotalRows(),
		},
	})
	return nil
}

func (h CartHandlerImpl) Remove(ctx echo.Context) error {
	userId := ctx.Get("user_id").(float64)
	cartId, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	cart, err := h.SvcCart.RemoveProduct(uint(userId), uint(cartId))

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusNoContent, "Success", cart)
	return nil
}

func (h CartHandlerImpl) Add(ctx echo.Context) error {
	var request dto.CartAddProductRequest

	if err := ctx.Bind(&request); err != nil {
		response.Err(ctx, http.StatusBadRequest, fmt.Errorf(err.Error()))
		return err
	}

	userId := ctx.Get("user_id").(float64)

	cart, err := h.SvcCart.AddProduct(uint(userId), &request)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", cart)
	return nil
}

func (h CartHandlerImpl) Edit(ctx echo.Context) error {
	userId := ctx.Get("user_id").(float64)

	var request dto.CartEditProductRequest

	if err := ctx.Bind(&request); err != nil {
		response.Err(ctx, http.StatusBadRequest, fmt.Errorf(err.Error()))
		return err
	}

	cart, err := h.SvcCart.UpdateQuantity(uint(userId), request.Id, request.Quantity)

	if err != nil {
		response.Err(ctx, http.StatusBadRequest, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", cart)
	return nil
}
