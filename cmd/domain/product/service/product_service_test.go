package service

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
	"toko/cmd/domain/product/dto"
	"toko/cmd/domain/product/entity"
	"toko/cmd/domain/product/repository"
	"toko/pkg/database"
)

var productRepository = repository.ProductRepositoryMock{Mock: mock.Mock{}}
var productService = ProductServiceImpl{RepoProduct: &productRepository}

func TestProductService_GetProducts(t *testing.T) {
	t.Run("GetProductsFail", func(t *testing.T) {
		// program mock
		productRepository.Mock.On("FindAll").Return(nil, errors.New("products not found")).Once()

		e := echo.New()
		r := httptest.NewRequest("POST", "/product", nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		pagination := database.NewPagination(ctx)

		product, err := productService.GetProducts(pagination)
		assert.Nil(t, product)
		assert.NotNil(t, err)
	})

	t.Run("GetProductsSuccess", func(t *testing.T) {
		product := entity.Product{
			ID:          1,
			Name:        "Product 1",
			Description: "Ini Product 1",
			Slug:        "product-1",
			Stock:       10,
			Price:       20000,
		}
		productList := entity.ProductList{&product}
		// program mock
		productRepository.Mock.On("FindAll").Return(productList, nil).Once()

		e := echo.New()
		r := httptest.NewRequest("GET", "/product", nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		pagination := database.NewPagination(ctx)

		products, err := productService.GetProducts(pagination)
		assert.NotNil(t, products)
		assert.Nil(t, err)
		assert.IsType(t, dto.ProductListResponse{}, *products)
		assert.Equal(t, product.ID, (*products)[0].ID)
	})
}

func TestProductService_GetProductById(t *testing.T) {
	product := entity.Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Ini Product 1",
		Slug:        "product-1",
		Stock:       10,
		Price:       20000,
	}

	t.Run("GetProductByIdFail", func(t *testing.T) {
		// program mock
		productRepository.Mock.On("Find", product.ID).Return(nil, errors.New("product not found")).Once()

		product, err := productService.GetProductById(product.ID)
		assert.Nil(t, product)
		assert.NotNil(t, err)
	})

	t.Run("GetProductByIdSuccess", func(t *testing.T) {
		// program mock
		productRepository.Mock.On("Find", product.ID).Return(product, nil).Once()

		productRes, err := productService.GetProductById(product.ID)
		assert.NotNil(t, productRes)
		assert.Nil(t, err)
		assert.IsType(t, dto.ProductResponse{}, *productRes)
		assert.Equal(t, product.ID, productRes.ID)
	})
}

func TestProductService_GetProductBySlug(t *testing.T) {
	product := entity.Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Ini Product 1",
		Slug:        "product-1",
		Stock:       10,
		Price:       20000,
	}

	t.Run("GetProductBySlugFail", func(t *testing.T) {
		// program mock
		productRepository.Mock.On("FindBySlug", product.Slug).Return(nil, errors.New("product not found")).Once()

		product, err := productService.GetProductBySlug(product.Slug)
		assert.Nil(t, product)
		assert.NotNil(t, err)
	})

	t.Run("GetProductBySlugSuccess", func(t *testing.T) {
		// program mock
		productRepository.Mock.On("FindBySlug", product.Slug).Return(product, nil).Once()

		productRes, err := productService.GetProductBySlug(product.Slug)
		assert.NotNil(t, productRes)
		assert.Nil(t, err)
		assert.IsType(t, dto.ProductResponse{}, *productRes)
		assert.Equal(t, product.ID, productRes.ID)
	})
}

func TestProductService_Store(t *testing.T) {
	product := entity.Product{
		ID:          1,
		Name:        "Product 1",
		Description: "Ini Product 1",
		Slug:        "product-1",
		Stock:       10,
		Price:       20000,
	}

	productReq := &dto.ProductStoreRequest{
		Name:        "Product 1",
		Description: "Ini Product 1",
		Stock:       10,
		Price:       20000,
	}

	t.Run("StoreProductFail", func(t *testing.T) {
		// program mock
		productRepository.Mock.On("Insert", mock.Anything).Return(nil, errors.New("insert product fail")).Once()

		product, err := productService.Store(productReq)

		assert.Nil(t, product)
		assert.NotNil(t, err)
	})

	t.Run("StoreProductSuccess", func(t *testing.T) {
		// program mock
		productRepository.Mock.On("Insert", mock.Anything).Return(product, nil).Once()

		productRes, err := productService.Store(productReq)

		assert.NotNil(t, productRes)
		assert.Nil(t, err)
		assert.IsType(t, dto.ProductResponse{}, *productRes)
		assert.Equal(t, product.ID, productRes.ID)
	})
}
