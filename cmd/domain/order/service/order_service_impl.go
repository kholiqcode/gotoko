package service

import (
	"database/sql"
	"fmt"
	"github.com/labstack/gommon/random"
	"github.com/rs/zerolog/log"
	"strings"
	_rCart "toko/cmd/domain/cart/repository"
	"toko/cmd/domain/order/dto"
	"toko/cmd/domain/order/entity"
	_rOrder "toko/cmd/domain/order/repository"
	"toko/pkg/database"
)

type OrderServiceImpl struct {
	RepoOrder _rOrder.OrderRepository
	RepoCart  _rCart.CartRepository
}

func (s OrderServiceImpl) GetOrders(pagination *database.Pagination, userId uint) (*dto.OrderListResponse, error) {
	orders, err := s.RepoOrder.FindAll(pagination, userId)
	if err != nil {
		log.Err(err).Msg("Error fetch orders from DB")
		return nil, err
	}
	orderResp := dto.CreateOrderListResponse(orders)
	return &orderResp, nil
}

func (s OrderServiceImpl) Create(request *dto.OrderRequest, userId uint) (*dto.OrderResponse, error) {
	carts, err := s.RepoCart.FindByUser(userId, request.Carts)

	if err != nil {
		log.Err(err).Msg("your cart is empty")
		return nil, err
	}

	var orderDetails []*entity.OrderDetail
	var total float64

	for _, cart := range *carts {
		total += cart.SubTotal
		orderDetails = append(orderDetails, &entity.OrderDetail{
			ProductID:          cart.Product.ID,
			ProductName:        cart.Product.Name,
			ProductImage:       cart.Product.FeaturedImage,
			ProductDescription: cart.Product.Description,
			ProductPrice:       cart.Product.Price,
			Quantity:           cart.Quantity,
			SubTotal:           cart.SubTotal,
		})
	}

	invoiceNo := fmt.Sprintf("INV-%s", strings.ToUpper(random.String(16)))
	orderRepo, err := s.RepoOrder.Insert(&entity.Order{
		UserID:        userId,
		InvoiceNo:     invoiceNo,
		PaymentMethod: request.PaymentMethod,
		CustomerName:  "hello",
		CustomerEmail: "hello",
		Note:          sql.NullString{String: request.Note, Valid: true},
		Status:        1,
		Total:         total,
		OrderDetail:   orderDetails,
	})
	if err != nil {
		log.Err(err).Msg("Error insert new order to DB")
		return nil, err
	}

	orderResp := dto.CreateOrderResponse(orderRepo)

	return &orderResp, nil
}
