package dto

import (
	"time"
	"toko/cmd/domain/product/entity"
)

type ProductResponse struct {
	ID              uint                      `json:"id"`
	Name            string                    `json:"name"`
	FeaturedImage   string                    `json:"featured_image"`
	Description     string                    `json:"description"`
	Slug            string                    `json:"slug"`
	Stock           uint                      `json:"stock"`
	Price           float64                   `json:"price"`
	ProductGallery  []ProductGalleryResponse  `json:"product_galleries"`
	ProductCategory []ProductCategoryResponse `json:"product_categories"`
	CreatedAt       time.Time                 `json:"created_at,omitempty"`
	UpdatedAt       time.Time                 `json:"updated_at,omitempty"`
}

type ProductListResponse []*ProductResponse

type ProductGalleryResponse struct {
	ID             uint   `json:"id"`
	Path           string `json:"path"`
	AltTitle       string `json:"alt_title"`
	AltDescription string `json:"alt_description"`
}

type ProductCategoryResponse struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Slug           string `json:"slug"`
	AltTitle       string `json:"alt_title"`
	AltDescription string `json:"alt_description"`
}

func CreateProductResponse(product *entity.Product) ProductResponse {

	var categoryResp []ProductCategoryResponse
	for _, c := range product.ProductCategory {
		categoryResp = append(categoryResp, ProductCategoryResponse{
			ID:   c.Category.ID,
			Name: c.Category.Name,
			Slug: c.Category.Slug,
		})
	}

	var galleryResp []ProductGalleryResponse
	var featuredImage string
	if len(product.ProductGallery) > 0 {
		featuredImage = product.ProductGallery[0].Path
		for _, c := range product.ProductGallery {
			galleryResp = append(galleryResp, ProductGalleryResponse{
				ID:             c.ID,
				Path:           c.Path,
				AltTitle:       c.AltTitle.String,
				AltDescription: c.AltDescription.String,
			})
		}
	}

	productResp := ProductResponse{
		ID:              product.ID,
		Name:            product.Name,
		FeaturedImage:   featuredImage,
		Description:     product.Description,
		Slug:            product.Slug,
		Stock:           product.Stock,
		Price:           product.Price,
		ProductCategory: categoryResp,
		ProductGallery:  galleryResp,
		CreatedAt:       product.CreatedAt,
		UpdatedAt:       product.UpdatedAt,
	}
	return productResp
}

func CreateProductListResponse(products *entity.ProductList) ProductListResponse {
	productResp := ProductListResponse{}
	for _, p := range *products {
		user := CreateProductResponse(p)
		productResp = append(productResp, &user)
	}
	return productResp
}
