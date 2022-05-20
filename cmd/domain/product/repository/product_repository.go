package repository

import (
	"toko/cmd/domain/product/entity"
	"toko/pkg/database"
)

type ProductRepository interface {
	FindAll(pagination *database.Pagination) (*entity.ProductList, error)
	FindAllCategory(pagination *database.Pagination) (*entity.CategoryList, error)
	Find(productID uint) (*entity.Product, error)
	FindBySlug(slug string) (*entity.Product, error)
	Insert(product *entity.Product) (*entity.Product, error)
	InsertCategory(category *entity.Category) (*entity.Category, error)
}
