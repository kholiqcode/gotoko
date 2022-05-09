package repository

import (
	"toko/cmd/domain/product/entity"
	"toko/pkg/database"
)

type ProductRepository interface {
	FindAll(pagination *database.Pagination) (*entity.ProductList, error)
	Find(productID uint) (*entity.Product, error)
	FindBySlug(slug string) (*entity.Product, error)
	Insert(product *entity.Product) (*entity.Product, error)
}
