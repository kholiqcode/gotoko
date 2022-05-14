package service

import (
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
	"toko/cmd/domain/product/dto"
	"toko/cmd/domain/product/entity"
	"toko/cmd/domain/product/repository"
	"toko/pkg/database"
)

type ProductServiceImpl struct {
	Repo repository.ProductRepository
}

func (s ProductServiceImpl) GetProducts(pagination *database.Pagination) (*dto.ProductListResponse, error) {
	products, err := s.Repo.FindAll(pagination)
	if err != nil {
		log.Err(err).Msg("Error fetch products from DB")
		return nil, err
	}
	usersResp := dto.CreateProductListResponse(products)
	return &usersResp, nil
}

func (s ProductServiceImpl) GetProductById(productId uint) (*dto.ProductResponse, error) {
	product, err := s.Repo.Find(productId)
	if err != nil {
		log.Err(err).Msg("Error fetch product from DB")
		return nil, err
	}
	productResp := dto.CreateProductResponse(product)
	return &productResp, nil
}

func (s ProductServiceImpl) GetProductBySlug(slug string) (*dto.ProductResponse, error) {
	product, err := s.Repo.FindBySlug(slug)
	if err != nil {
		log.Err(err).Msg("Error fetch product from DB")
		return nil, err
	}
	productResp := dto.CreateProductResponse(product)
	return &productResp, nil
}

func (s ProductServiceImpl) Store(request *dto.ProductStoreRequest) (*dto.ProductResponse, error) {
	var productCategory []entity.ProductCategory
	for _, c := range request.Categories {
		productCategory = append(productCategory, entity.ProductCategory{
			CategoryID: c,
		})
	}

	slug := slug.Make(request.Name)
	productRepo, err := s.Repo.Insert(&entity.Product{
		Name:            request.Name,
		Description:     request.Description,
		Stock:           request.Stock,
		Price:           request.Price,
		Slug:            slug,
		ProductCategory: productCategory,
	})
	if err != nil {
		log.Err(err).Msg("Error insert product to DB")
		return nil, err
	}

	productResp := dto.CreateProductResponse(productRepo)

	return &productResp, nil

}
