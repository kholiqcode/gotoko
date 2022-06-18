package product

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
	"toko/cmd/domain/product/dto"
	"toko/cmd/domain/product/entity"
	"toko/cmd/domain/product/service"
)

var productService = &service.ProductServiceMock{Mock: mock.Mock{}}
var productHandler = ProductHandlerImpl{SvcProduct: productService}

func TestProductHandler_Get(t *testing.T) {
	product := entity.Product{
		ID:          1,
		Name:        "Product 1",
		Slug:        "product-1",
		Description: "deskripsi product 1",
		Price:       30000,
		Stock:       12,
	}
	productList := entity.ProductList{&product}

	t.Run("GetProductSuccess", func(t *testing.T) {
		//program mock
		productService.Mock.On("GetProducts", mock.Anything).Return(dto.CreateProductListResponse(&productList), nil).Once()

		dataResponse := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
			Result  struct {
				Products struct {
					Data dto.ProductListResponse `json:"data"`
					Page int                     `json:"page"`
				} `json:"products"`
			} `json:"result"`
		}{}

		e := echo.New()
		r := httptest.NewRequest("GET", "/product", nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)
		productHandler.Get(ctx)
		bodyRes, _ := ioutil.ReadAll(w.Result().Body)
		var err = json.Unmarshal(bodyRes, &dataResponse)
		assert.NoError(t, err)

		assert.Equal(t, 200, w.Result().StatusCode)
		assert.EqualValues(t, dataResponse.Message, "Success")
		assert.EqualValues(t, dataResponse.Result.Products.Data[0].ID, product.ID)
		assert.EqualValues(t, dataResponse.Result.Products.Page, product.ID)
	})

	t.Run("GetProductFail", func(t *testing.T) {
		//program mock
		productService.Mock.On("GetProducts", mock.Anything).Return(nil, errors.New("fail on get service")).Once()

		dataResponse := struct {
			Code    uint   `json:"code"`
			Message string `json:"message"`
			Result  struct {
				Products struct {
					Data dto.ProductListResponse `json:"data"`
					Page int                     `json:"page"`
				} `json:"products"`
			} `json:"result"`
		}{}

		e := echo.New()
		r := httptest.NewRequest("GET", "/product", nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)
		productHandler.Get(ctx)
		bodyRes, _ := ioutil.ReadAll(w.Result().Body)
		var err = json.Unmarshal(bodyRes, &dataResponse)
		assert.NoError(t, err)

		assert.Equal(t, 400, w.Result().StatusCode)
		assert.NotEqualValues(t, dataResponse.Message, "Success")
	})
}

func TestProductHandler_Detail(t *testing.T) {
	product := entity.Product{
		ID:          1,
		Name:        "Product 1",
		Slug:        "product-1",
		Description: "deskripsi product 1",
		Price:       30000,
		Stock:       12,
	}

	t.Run("GetProductDetailSuccess", func(t *testing.T) {
		//program mock
		productService.Mock.On("GetProductBySlug", mock.Anything).Return(dto.CreateProductResponse(&product), nil).Once()

		dataResponse := struct {
			Code    uint                `json:"code"`
			Message string              `json:"message"`
			Result  dto.ProductResponse `json:"result"`
		}{}

		e := echo.New()
		r := httptest.NewRequest("GET", "/product/"+product.Slug, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)
		productHandler.Detail(ctx)
		bodyRes, _ := ioutil.ReadAll(w.Result().Body)
		var err = json.Unmarshal(bodyRes, &dataResponse)
		assert.NoError(t, err)

		assert.Equal(t, 200, w.Result().StatusCode)
		assert.EqualValues(t, dataResponse.Message, "Success")
		assert.EqualValues(t, dataResponse.Result.ID, product.ID)
	})

	t.Run("GetProductDetailFail", func(t *testing.T) {
		//program mock
		productService.Mock.On("GetProductBySlug", mock.Anything).Return(nil, errors.New("fail on get service")).Once()

		dataResponse := struct {
			Code    uint                `json:"code"`
			Message string              `json:"message"`
			Data    dto.ProductResponse `json:"data"`
		}{}

		e := echo.New()
		r := httptest.NewRequest("GET", "/product/"+product.Slug, nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)
		productHandler.Detail(ctx)
		bodyRes, _ := ioutil.ReadAll(w.Result().Body)
		var err = json.Unmarshal(bodyRes, &dataResponse)
		assert.NoError(t, err)

		assert.Equal(t, 400, w.Result().StatusCode)
		assert.NotEqualValues(t, dataResponse.Message, "Success")
	})
}

func TestProductHandler_Create(t *testing.T) {
	res := dto.ProductResponse{
		ID:          1,
		Name:        "Product 1",
		Slug:        "product-1",
		Description: "deskripsi product 1",
		Price:       30000,
		Stock:       12,
	}

	t.Run("CreateProductSuccess", func(t *testing.T) {
		//program mock
		productService.Mock.On("Store", mock.Anything).Return(res, nil).Once()

		dataResponse := struct {
			Code    uint                `json:"code"`
			Message string              `json:"message"`
			Result  dto.ProductResponse `json:"result"`
		}{}

		e := echo.New()
		bodyReq := make(map[string]interface{})
		bodyReq["name"] = "Product 1"
		bodyReq["description"] = "deskripsi product 1"
		bodyReq["price"] = 23000
		bodyReq["stock"] = 11
		bodyJson, _ := json.Marshal(bodyReq)
		r := httptest.NewRequest("POST", "/product", bytes.NewReader(bodyJson))
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)
		productHandler.Create(ctx)
		bodyRes, _ := ioutil.ReadAll(w.Result().Body)
		var err = json.Unmarshal(bodyRes, &dataResponse)
		assert.NoError(t, err)
		defer r.Body.Close()

		assert.Equal(t, 201, w.Result().StatusCode)
		assert.EqualValues(t, dataResponse.Message, "Success", "Message Equal")
		assert.EqualValues(t, dataResponse.Result.ID, res.ID, "Product Equal")
	})

	t.Run("CreateProductFail", func(t *testing.T) {
		//program mock
		productService.Mock.On("Store", mock.Anything).Return(nil, errors.New("fail on get service")).Once()

		dataResponse := struct {
			Code    uint                `json:"code"`
			Message string              `json:"message"`
			Result  dto.ProductResponse `json:"result"`
		}{}

		e := echo.New()
		bodyReq := make(map[string]interface{})
		bodyReq["name"] = "Product 1"
		bodyReq["description"] = "deskripsi product 1"
		bodyReq["price"] = 23000
		bodyReq["stock"] = 11
		bodyJson, _ := json.Marshal(bodyReq)
		r := httptest.NewRequest("POST", "/product", bytes.NewReader(bodyJson))
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)
		productHandler.Create(ctx)
		bodyRes, _ := ioutil.ReadAll(w.Result().Body)
		var err = json.Unmarshal(bodyRes, &dataResponse)
		assert.NoError(t, err)
		defer r.Body.Close()

		assert.Equal(t, 400, w.Result().StatusCode)
		assert.NotEqualValues(t, dataResponse.Message, "Success", "Message Not Equal")
		assert.NotEqualValues(t, dataResponse.Result.ID, res.ID, "Product ID Not Equal")
	})
}
