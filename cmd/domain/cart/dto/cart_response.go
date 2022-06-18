package dto

import (
	"time"
	"toko/cmd/domain/cart/entity"
)

type CartResponse struct {
	ID        uint            `json:"id"`
	Quantity  uint            `json:"quantity"`
	SubTotal  float64         `json:"subtotal,omitempty"`
	Product   ProductResponse `json:"product_detail,omitempty"`
	CreatedAt time.Time       `json:"created_at,omitempty"`
	UpdatedAt time.Time       `json:"updated_at,omitempty"`
}

type CartListResponse []*CartResponse

type ProductResponse struct {
	ID            uint    `json:"id,omitempty"`
	Name          string  `json:"name,omitempty"`
	FeaturedImage string  `json:"featured_image,omitempty"`
	Description   string  `json:"description,omitempty"`
	Slug          string  `json:"slug,omitempty"`
	Stock         uint    `json:"stock,omitempty"`
	Price         float64 `json:"price,omitempty"`
}

func CreateCartResponse(cart *entity.Cart) CartResponse {
	var featuredImage string
	var productResponse ProductResponse
	if len(cart.Product.ProductGallery) > 0 {
		featuredImage = cart.Product.ProductGallery[0].Path
	}

	productResponse = ProductResponse{
		ID:            cart.Product.ID,
		Name:          cart.Product.Name,
		FeaturedImage: featuredImage,
		Description:   cart.Product.Description,
		Slug:          cart.Product.Slug,
		Stock:         cart.Product.Stock,
		Price:         cart.Product.Price,
	}

	cartResp := CartResponse{
		ID:        cart.ID,
		Quantity:  cart.Quantity,
		SubTotal:  cart.SubTotal,
		Product:   productResponse,
		CreatedAt: cart.CreatedAt,
		UpdatedAt: cart.UpdatedAt,
	}

	return cartResp
}

func CreateCartListResponse(carts *entity.CartList) CartListResponse {
	cartResp := CartListResponse{}
	for _, p := range *carts {
		cart := CreateCartResponse(p)
		cartResp = append(cartResp, &cart)
	}
	return cartResp
}
