package service

import (
	"toko/cmd/domain/product/dto"
	"toko/pkg/database"
)

type ProductService interface {
	GetProducts(pagination *database.Pagination) (*dto.ProductListResponse, error)
	GetCategories(pagination *database.Pagination) (*dto.CategoryListResponse, error)
	GetProductById(productId uint) (*dto.ProductResponse, error)
	GetProductBySlug(slug string) (*dto.ProductResponse, error)
	Store(request *dto.ProductStoreRequest) (*dto.ProductResponse, error)
	StoreCategory(request *dto.CategoryStoreRequest) (*dto.CategoryResponse, error)
}
