package dto

type CartAddProductRequest struct {
	Product  uint `json:"product" form:"product"`
	Quantity uint `json:"quantity" form:"quantity"`
}

type CartEditProductRequest struct {
	Id       uint `json:"id" form:"id"`
	Quantity uint `json:"quantity" form:"quantity"`
}
