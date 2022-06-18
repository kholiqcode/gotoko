package repository

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"toko/cmd/domain/product/entity"
	"toko/pkg/database"
)

type ProductRepositoryMock struct {
	Mock mock.Mock
}

func (r *ProductRepositoryMock) FindAllCategory(pagination *database.Pagination) (*entity.CategoryList, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ProductRepositoryMock) InsertCategory(category *entity.Category) (*entity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ProductRepositoryMock) FindAll(pagination *database.Pagination) (*entity.ProductList, error) {
	arguments := r.Mock.Called()
	if arguments.Get(0) == nil {
		return nil, errors.New("not found arguments")
	} else {
		products := arguments.Get(0).(entity.ProductList)
		return &products, nil
	}
}
func (r *ProductRepositoryMock) Find(productID uint) (*entity.Product, error) {
	arguments := r.Mock.Called(productID)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found productId")
	}
	product := arguments.Get(0).(entity.Product)
	return &product, nil
}
func (r *ProductRepositoryMock) FindBySlug(slug string) (*entity.Product, error) {
	arguments := r.Mock.Called(slug)
	if arguments.Get(0) == nil {
		return nil, errors.New("not found slug")
	}
	product := arguments.Get(0).(entity.Product)
	return &product, nil
}
func (r *ProductRepositoryMock) Insert(product *entity.Product) (*entity.Product, error) {
	arguments := r.Mock.Called(product)
	if arguments.Get(0) == nil {
		return nil, errors.New("err")
	} else {
		product := arguments.Get(0).(entity.Product)
		return &product, nil
	}
}
