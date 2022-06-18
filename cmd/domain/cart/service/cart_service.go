package service

import (
	"toko/cmd/domain/cart/dto"
	"toko/pkg/database"
)

type CartService interface {
	GetCarts(pagination *database.Pagination, userId uint) (*dto.CartListResponse, error)
	AddProduct(userId uint, cart *dto.CartAddProductRequest) (*dto.CartResponse, error)
	RemoveProduct(userId uint, cartId uint) (*dto.CartResponse, error)
	UpdateQuantity(userId uint, cartId uint, newQty uint) (*dto.CartResponse, error)
}
