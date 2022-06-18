package order

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"toko/cmd/domain/order/dto"
	_sOrder "toko/cmd/domain/order/service"
	"toko/internal/protocol/http/response"
	"toko/pkg/database"
)

type OrderHandlerImpl struct {
	SvcOrder _sOrder.OrderService
}

func (h OrderHandlerImpl) Get(ctx echo.Context) error {
	pagination := database.NewPagination(ctx)
	userId := ctx.Get("user_id").(float64)

	products, err := h.SvcOrder.GetOrders(pagination, uint(userId))

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"orders": map[string]interface{}{
			"data":       products,
			"sort":       pagination.GetSort(),
			"page":       pagination.GetPage(),
			"page_size":  pagination.GetLimit(),
			"total_page": pagination.GetTotalPage(),
			"total_rows": pagination.GetTotalRows(),
		},
	})
	return nil
}

func (h OrderHandlerImpl) Create(ctx echo.Context) error {
	userId := ctx.Get("user_id").(float64)
	var request dto.OrderRequest

	if err := ctx.Bind(&request); err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	order, err := h.SvcOrder.Create(&request, uint(userId))

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", order)
	return nil
}
