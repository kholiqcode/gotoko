package dto

import "mime/multipart"

type ProductStoreRequest struct {
	Name        string                  `json:"name" form:"name"`
	Description string                  `json:"description" form:"description"`
	Stock       uint                    `json:"stock" form:"stock"`
	Price       float64                 `json:"price" form:"price"`
	Categories  []uint                  `json:"categories" form:"categories"`
	Images      []*multipart.FileHeader `json:"images" form:"images"`
}
