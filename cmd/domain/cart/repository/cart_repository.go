package repository

import (
	"toko/cmd/domain/cart/entity"
	"toko/pkg/database"
)

type CartRepository interface {
	FindAll(pagination *database.Pagination, userId uint) (*entity.CartList, error)
	FindByUser(userId uint, cartsId []uint) (*entity.CartList, error)
	FindByProductId(userId uint, productId uint) (*entity.CartList, error)
	Insert(cart *entity.Cart) (*entity.Cart, error)
	UpdateQuantity(userId uint, cartId uint, newQty uint) (*entity.Cart, error)
	Delete(userId uint, cartId uint) (*entity.Cart, error)
}
