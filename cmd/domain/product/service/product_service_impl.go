package service

import (
	"github.com/rs/zerolog/log"
	"toko/cmd/domain/product/dto"
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
	//TODO implement me
	panic("implement me")
}

func (s ProductServiceImpl) Store(request *dto.ProductStoreRequest) (*dto.ProductResponse, error) {
	//TODO implement me
	panic("implement me")
}
