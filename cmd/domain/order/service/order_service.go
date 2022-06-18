package service

import (
	"toko/cmd/domain/order/dto"
	"toko/pkg/database"
)

type OrderService interface {
	GetOrders(pagination *database.Pagination, userId uint) (*dto.OrderListResponse, error)
	Create(request *dto.OrderRequest, userId uint) (*dto.OrderResponse, error)
}
