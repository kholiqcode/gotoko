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

func (r ProductRepositoryImpl) FindAllCategory(pagination *database.Pagination) (*entity.CategoryList, error) {
	var categories entity.CategoryList

	if e := r.Db.Debug().Scopes(database.Paginate(categories, pagination, r.Db)).Find(&categories).Error; e != nil {
		return nil, e
	}

	return &categories, nil
}

func (r ProductRepositoryImpl) Find(productID uint) (*entity.Product, error) {
	var product entity.Product

	if e := r.Db.Debug().Preload("ProductCategory.Category").Preload("ProductGallery").Preload(clause.Associations).First(&product, productID).Error; e != nil {
		return nil, e
	}

	return &product, nil
}

func (r ProductRepositoryImpl) FindBySlug(slug string) (*entity.Product, error) {
	var product entity.Product

	if e := r.Db.Debug().Preload("ProductCategory.Category").Preload("ProductGallery").Preload(clause.Associations).First(&product, "slug = ?", slug).Error; e != nil {
		return nil, e
	}

	return &product, nil
}

func (r ProductRepositoryImpl) Insert(product *entity.Product) (*entity.Product, error) {
	if e := r.Db.Debug().Create(&product).Preload("ProductCategory.Category").Error; e != nil {
		return nil, e
	}
	return product, nil
}

func (r ProductRepositoryImpl) InsertCategory(category *entity.Category) (*entity.Category, error) {
	if e := r.Db.Debug().Create(&category).Error; e != nil {
		return nil, e
	}
	return category, nil
}
