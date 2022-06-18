package service

import (
	"errors"
	"github.com/rs/zerolog/log"
	"toko/cmd/domain/cart/dto"
	"toko/cmd/domain/cart/entity"
	_rCart "toko/cmd/domain/cart/repository"
	_rProduct "toko/cmd/domain/product/repository"
	"toko/pkg/database"
)

type CartServiceImpl struct {
	RepoCart    _rCart.CartRepository
	RepoProduct _rProduct.ProductRepository
}

func (s CartServiceImpl) GetCarts(pagination *database.Pagination, userId uint) (*dto.CartListResponse, error) {
	carts, err := s.RepoCart.FindAll(pagination, userId)
	if err != nil {
		log.Err(err).Msg("Error fetch carts from DB")
		return nil, err
	}
	cartResp := dto.CreateCartListResponse(carts)
	return &cartResp, nil
}

func (s CartServiceImpl) AddProduct(userId uint, request *dto.CartAddProductRequest) (*dto.CartResponse, error) {
	product, err := s.RepoProduct.Find(request.Product)
	if err != nil {
		log.Err(err).Msg("Product not found")
		return nil, errors.New("product not found")
	}

	if request.Quantity > product.Stock {
		log.Err(err).Msg("product out of stock")
		return nil, errors.New("product out of stock")
	}

	cart, err := s.RepoCart.Insert(&entity.Cart{
		UserID:    userId,
		ProductID: request.Product,
		Quantity:  request.Quantity,
	})
	if err != nil {
		log.Err(err).Msg("Error add product to cart")
		return nil, err
	}

	cartResp := dto.CreateCartResponse(cart)
	return &cartResp, nil
}

func (s CartServiceImpl) UpdateQuantity(userId uint, cartId uint, newQty uint) (*dto.CartResponse, error) {
	cart, err := s.RepoCart.UpdateQuantity(userId, cartId, newQty)
	if err != nil {
		log.Err(err).Msg("can't update quantity")
		return nil, errors.New("cart not found")
	}
	cartResp := dto.CreateCartResponse(cart)
	return &cartResp, nil
}

func (s CartServiceImpl) RemoveProduct(userId uint, cartId uint) (*dto.CartResponse, error) {
	cart, err := s.RepoCart.Delete(userId, cartId)
	if err != nil {
		log.Err(err).Msg("can't delete cart")
		return nil, err
	}
	cartResp := dto.CreateCartResponse(cart)
	return &cartResp, nil
}
