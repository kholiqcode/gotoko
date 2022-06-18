package repository

import (
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"toko/cmd/domain/cart/entity"
	"toko/pkg/database"
)

type CartRepositoryImpl struct {
	DbMysql *gorm.DB
}

func (r CartRepositoryImpl) FindAll(pagination *database.Pagination, userId uint) (*entity.CartList, error) {
	var carts entity.CartList

	if e := r.DbMysql.Debug().Where("user_id = ?", userId).Scopes(database.Paginate(carts, pagination, r.DbMysql)).Preload("Product").Preload(clause.Associations).Find(&carts).Error; e != nil {
		return nil, e
	}

	return &carts, nil
}
func (r CartRepositoryImpl) FindByUser(userId uint, cartsId []uint) (*entity.CartList, error) {
	var carts entity.CartList

	if e := r.DbMysql.Debug().Where("user_id = ?", userId).Preload("Product").Preload(clause.Associations).Find(&carts, cartsId).Error; e != nil {
		return nil, e
	}

	return &carts, nil
}
func (r CartRepositoryImpl) FindByProductId(userId uint, productId uint) (*entity.CartList, error) {
	var carts entity.CartList

	if e := r.DbMysql.Debug().Where("user_id = ?", userId).Where("product_id = ?", productId).First(&carts).Error; e != nil {
		log.Err(e).Msg("Error find cart to DB")
		return nil, e
	}

	return &carts, nil
}

func (r CartRepositoryImpl) Insert(cart *entity.Cart) (*entity.Cart, error) {
	if e := r.DbMysql.Debug().Preload("Product").Where(entity.Cart{UserID: cart.UserID}).Where(entity.Cart{ProductID: cart.ProductID}).Assign(map[string]interface{}{"quantity": gorm.Expr("quantity + ?", cart.Quantity)}).FirstOrCreate(&cart).Error; e != nil {
		log.Err(e).Msg("Error insert cart to DB")
		return nil, e
	}
	return cart, nil
}

func (r CartRepositoryImpl) UpdateQuantity(userId uint, cartId uint, newQty uint) (*entity.Cart, error) {
	cart := &entity.Cart{}
	if e := r.DbMysql.Debug().Preload("Product").Where("user_id = ?", userId).First(&cart, cartId).Error; e != nil {
		log.Err(e).Msg("Error update quantity cart")
		return nil, e
	}

	cart.Quantity = newQty
	r.DbMysql.Save(cart)

	return cart, nil
}

func (r CartRepositoryImpl) Delete(userId uint, cartId uint) (*entity.Cart, error) {
	cart := &entity.Cart{}
	if e := r.DbMysql.Debug().Where("user_id = ?", userId).Delete(cart, cartId).Error; e != nil {
		return nil, e
	}
	return cart, nil
}
