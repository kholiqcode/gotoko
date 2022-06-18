package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"toko/cmd/domain/order/entity"
	"toko/pkg/database"
)

type OrderRepositoryImpl struct {
	DbMysql *gorm.DB
}

func (r OrderRepositoryImpl) FindAll(pagination *database.Pagination, userId uint) (*entity.OrderList, error) {
	var orders entity.OrderList

	if e := r.DbMysql.Debug().Where("user_id = ?", userId).Scopes(database.Paginate(orders, pagination, r.DbMysql)).Preload("OrderDetail.Product.ProductGallery").Preload("User").Preload(clause.Associations).Find(&orders).Error; e != nil {
		return nil, e
	}

	return &orders, nil
}

func (r OrderRepositoryImpl) Insert(order *entity.Order) (*entity.Order, error) {
	if e := r.DbMysql.Debug().Preload("OrderDetail").Create(&order).Error; e != nil {
		return nil, e
	}
	return order, nil
}
