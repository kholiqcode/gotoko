package service

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"toko/cmd/domain/product/dto"
	"toko/pkg/database"
)

type ProductServiceMock struct {
	Mock mock.Mock
}

func (m ProductServiceMock) GetCategories(pagination *database.Pagination) (*dto.CategoryListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m ProductServiceMock) StoreCategory(request *dto.CategoryStoreRequest) (*dto.CategoryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m ProductServiceMock) GetProducts(pagination *database.Pagination) (*dto.ProductListResponse, error) {
	arguments := m.Mock.Called(pagination)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		products := arguments.Get(0).(dto.ProductListResponse)
		return &products, nil
	}
}
func (m ProductServiceMock) GetProductById(productId uint) (*dto.ProductResponse, error) {
	arguments := m.Mock.Called(productId)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		product := arguments.Get(0).(dto.ProductResponse)
		return &product, nil
	}
}
func (m ProductServiceMock) GetProductBySlug(slug string) (*dto.ProductResponse, error) {
	arguments := m.Mock.Called(slug)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		product := arguments.Get(0).(dto.ProductResponse)
		return &product, nil
	}
}
func (m ProductServiceMock) Store(request *dto.ProductStoreRequest) (*dto.ProductResponse, error) {
	arguments := m.Mock.Called(request)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		product := arguments.Get(0).(dto.ProductResponse)
		return &product, nil
	}
}
