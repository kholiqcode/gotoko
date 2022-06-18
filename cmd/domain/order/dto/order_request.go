package dto

type OrderRequest struct {
	Carts         []uint `json:"carts" form:"carts"`
	PaymentMethod string `json:"payment_method" form:"payment_method"`
	Note          string `json:"note" form:"note"`
}
