package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"toko/cmd/domain/product/entity"
	"toko/pkg/database"
)

type ProductRepositoryImpl struct {
	Db *gorm.DB
}

func (r ProductRepositoryImpl) FindAll(pagination *database.Pagination) (*entity.ProductList, error) {
	var products entity.ProductList

	if e := r.Db.Debug().Scopes(database.Paginate(products, pagination, r.Db)).Preload("ProductCategory.Category").Preload("ProductGallery").Preload(clause.Associations).Find(&products).Error; e != nil {
		return nil, e
	}

	return &products, nil
}

func (r ProductRepositoryImpl) Find(productID uint) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r ProductRepositoryImpl) FindBySlug(slug string) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (r ProductRepositoryImpl) Insert(product *entity.Product) (*entity.Product, error) {
	//TODO implement me
	panic("implement me")
}
