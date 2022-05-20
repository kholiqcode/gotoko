package service

import (
	"database/sql"
	"github.com/gosimple/slug"
	"github.com/rs/zerolog/log"
	"toko/cmd/domain/product/dto"
	"toko/cmd/domain/product/entity"
	"toko/cmd/domain/product/repository"
	"toko/pkg/database"
)

type ProductServiceImpl struct {
	RepoProduct repository.ProductRepository
}

func (s ProductServiceImpl) GetProducts(pagination *database.Pagination) (*dto.ProductListResponse, error) {
	products, err := s.RepoProduct.FindAll(pagination)
	if err != nil {
		log.Err(err).Msg("Error fetch products from DB")
		return nil, err
	}
	usersResp := dto.CreateProductListResponse(products)
	return &usersResp, nil
}

func (s ProductServiceImpl) GetCategories(pagination *database.Pagination) (*dto.CategoryListResponse, error) {
	categories, err := s.RepoProduct.FindAllCategory(pagination)
	if err != nil {
		log.Err(err).Msg("Error fetch categories from DB")
		return nil, err
	}
	categoryResp := dto.CreateCategoryListResponse(categories)
	return &categoryResp, nil
}

func (s ProductServiceImpl) GetProductById(productId uint) (*dto.ProductResponse, error) {
	product, err := s.RepoProduct.Find(productId)
	if err != nil {
		log.Err(err).Msg("Error fetch product from DB")
		return nil, err
	}
	productResp := dto.CreateProductResponse(product)
	return &productResp, nil
}

func (s ProductServiceImpl) GetProductBySlug(slug string) (*dto.ProductResponse, error) {
	product, err := s.RepoProduct.FindBySlug(slug)
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
	productRepo, err := s.RepoProduct.Insert(&entity.Product{
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

func (s ProductServiceImpl) StoreCategory(request *dto.CategoryStoreRequest) (*dto.CategoryResponse, error) {
	slug := slug.Make(request.Name)
	categoryRepo, err := s.RepoProduct.InsertCategory(&entity.Category{
		Name:           request.Name,
		Slug:           slug,
		AltTitle:       sql.NullString{String: request.AltTitle, Valid: true},
		AltDescription: sql.NullString{String: request.AltDescription, Valid: true},
	})
	if err != nil {
		log.Err(err).Msg("Error insert category to DB")
		return nil, err
	}
	categoryResp := dto.CreateCategoryResponse(categoryRepo)
	return &categoryResp, nil
}
