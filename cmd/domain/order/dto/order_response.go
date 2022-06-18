package dto

import (
	"time"
	"toko/cmd/domain/order/entity"
)

type OrderResponse struct {
	ID            uint                   `json:"id"`
	InvoiceNo     string                 `json:"invoice_no,omitempty"`
	PaymentMethod string                 `json:"payment_method,omitempty"`
	CustomerName  string                 `json:"customer_name,omitempty"`
	CustomerEmail string                 `json:"customer_email,omitempty"`
	Note          string                 `json:"note,omitempty"`
	Status        uint                   `json:"status,omitempty"`
	Total         float64                `json:"total,omitempty"`
	OrderDetail   []*OrderDetailResponse `json:"order_detail,omitempty"`
	CreatedAt     time.Time              `json:"created_at,omitempty"`
	UpdatedAt     time.Time              `json:"updated_at,omitempty"`
}

type OrderListResponse []*OrderResponse

type OrderDetailResponse struct {
	ID                 uint      `json:"id"`
	ProductName        string    `json:"product_name,omitempty"`
	ProductImage       string    `json:"product_image,omitempty"`
	ProductDescription string    `json:"product_description,omitempty"`
	ProductPrice       float64   `json:"product_price,omitempty"`
	Quantity           uint      `json:"quantity"`
	SubTotal           float64   `json:"subtotal,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
}

func CreateOrderResponse(order *entity.Order) OrderResponse {
	var featuredImage string
	if len(order.OrderDetail) > 0 && len(order.OrderDetail[0].Product.ProductGallery) > 0 {
		featuredImage = order.OrderDetail[0].Product.ProductGallery[0].Path
	}

	var orderDetailResp []*OrderDetailResponse
	for _, c := range order.OrderDetail {
		orderDetailResp = append(orderDetailResp, &OrderDetailResponse{
			ID:                 c.ID,
			ProductName:        c.ProductName,
			ProductImage:       featuredImage,
			ProductDescription: c.ProductDescription,
			ProductPrice:       c.ProductPrice,
			Quantity:           c.Quantity,
			SubTotal:           c.SubTotal,
			CreatedAt:          c.CreatedAt,
			UpdatedAt:          c.UpdatedAt,
		})
	}
	orderResp := OrderResponse{
		ID:            order.ID,
		InvoiceNo:     order.InvoiceNo,
		PaymentMethod: order.PaymentMethod,
		CustomerName:  order.CustomerName,
		CustomerEmail: order.CustomerEmail,
		Note:          order.Note.String,
		Status:        order.Status,
		Total:         order.Total,
		OrderDetail:   orderDetailResp,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
	}

	return orderResp
}

func CreateOrderListResponse(orders *entity.OrderList) OrderListResponse {
	ordersResp := OrderListResponse{}
	for _, p := range *orders {
		order := CreateOrderResponse(p)
		ordersResp = append(ordersResp, &order)
	}
	return ordersResp
}
