package product

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"toko/cmd/domain/product/dto"
	"toko/cmd/domain/product/service"
	"toko/internal/protocol/http/response"
	"toko/pkg/database"
)

type ProductHandlerImpl struct {
	SvcProduct service.ProductService
}

func (h ProductHandlerImpl) CreateCategory(ctx echo.Context) error {
	var request dto.CategoryStoreRequest

	if err := ctx.Bind(&request); err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	category, err := h.SvcProduct.StoreCategory(&request)

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", category)
	return nil
}

func (h ProductHandlerImpl) Get(ctx echo.Context) error {
	pagination := database.NewPagination(ctx)

	products, err := h.SvcProduct.GetProducts(pagination)

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"products": map[string]interface{}{
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

func (h ProductHandlerImpl) GetCategory(ctx echo.Context) error {

	pagination := database.NewPagination(ctx)

	categories, err := h.SvcProduct.GetCategories(pagination)

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	response.Json(ctx, http.StatusOK, "Success", map[string]interface{}{
		"products": map[string]interface{}{
			"data":       categories,
			"sort":       pagination.GetSort(),
			"page":       pagination.GetPage(),
			"page_size":  pagination.GetLimit(),
			"total_page": pagination.GetTotalPage(),
			"total_rows": pagination.GetTotalRows(),
		},
	})
	return nil
}

func (h ProductHandlerImpl) Detail(ctx echo.Context) error {
	slug := ctx.Param("slug")

	product, err := h.SvcProduct.GetProductBySlug(slug)

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}
	response.Json(ctx, http.StatusOK, "Success", product)
	return nil
}

func (h ProductHandlerImpl) Create(ctx echo.Context) error {
	var productStoreDto dto.ProductStoreRequest

	if err := ctx.Bind(&productStoreDto); err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	product, err := h.SvcProduct.Store(&productStoreDto)

	if err != nil {
		response.Err(ctx, 400, err)
		return err
	}

	response.Json(ctx, http.StatusCreated, "Success", product)
	return nil
}
