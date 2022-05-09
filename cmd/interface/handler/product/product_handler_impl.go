package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"toko/cmd/domain/product/service"
	"toko/internal/protocol/http/response"
	"toko/pkg/database"
)

type ProductHandlerImpl struct {
	Svc service.ProductService
}

func (h ProductHandlerImpl) Get(ctx echo.Context) error {

	pagination := database.NewPagination(ctx)

	products, err := h.Svc.GetProducts(pagination)

	if err != nil {
		response.Err(ctx, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"users": map[string]interface{}{
			"data":      products,
			"sort":      pagination.GetSort(),
			"page":      pagination.GetPage(),
			"pageSize":  pagination.GetLimit(),
			"totalPage": pagination.GetTotalPage(),
			"totalRows": pagination.GetTotalRows(),
		},
	})
	return nil
}

func (h ProductHandlerImpl) Detail(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

func (h ProductHandlerImpl) Create(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
