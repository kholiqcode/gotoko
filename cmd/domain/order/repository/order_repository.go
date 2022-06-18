package repository

import (
	"toko/cmd/domain/order/entity"
	"toko/pkg/database"
)

type OrderRepository interface {
	FindAll(pagination *database.Pagination, userId uint) (*entity.OrderList, error)
	Insert(order *entity.Order) (*entity.Order, error)
}
