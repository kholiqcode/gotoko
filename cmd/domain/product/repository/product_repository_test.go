package repository

import (
	"database/sql/driver"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
	"time"
	entity "toko/cmd/domain/product/entity"
	"toko/pkg/database"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func InitConnection() (sqlmock.Sqlmock, ProductRepositoryImpl) {
	sqlMockDB, sqlMock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlMockDB,
		SkipInitializeWithVersion: true,
	}))

	productRepo := ProductRepositoryImpl{
		Db: gormDB,
	}
	return sqlMock, productRepo
}

func TestProductRepository_FindAll(t *testing.T) {
	sqlMock, productRepo := InitConnection()
	t.Run("FindAllError", func(t *testing.T) {
		sqlMock.ExpectQuery("SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL ORDER BY Id desc LIMIT 10").
			WillReturnError(errors.New("can't fetch to mock db"))

		e := echo.New()
		r := httptest.NewRequest("POST", "/product", nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		pagination := database.NewPagination(ctx)
		products, err := productRepo.FindAll(pagination)

		assert.NotNil(t, err)
		assert.Nil(t, products)
	})

	t.Run("FindAllSuccess", func(t *testing.T) {
		sqlMock, productRepo := InitConnection()
		sqlMock.ExpectQuery("SELECT * FROM `products` WHERE `products`.`deleted_at` IS NULL ORDER BY Id desc LIMIT 10").WillReturnRows(sqlmock.NewRows([]string{"id", "name", "slug", "description", "stock", "price", "created_at", "updated_at"}))

		e := echo.New()
		r := httptest.NewRequest("POST", "/product", nil)
		r.Header.Set("Content-Type", "application/json; charset=UTF-8")
		w := httptest.NewRecorder()
		ctx := e.NewContext(r, w)

		pagination := database.NewPagination(ctx)
		products, err := productRepo.FindAll(pagination)

		assert.Nil(t, err)
		assert.NotNil(t, products)
	})
}

func TestProductRepository_Find(t *testing.T) {
	sqlMock, productRepo := InitConnection()
	var (
		id          uint      = 1
		name        string    = "product 1"
		slug        string    = "product-1"
		description string    = "ini product 1"
		stock       uint      = 22
		price       float64   = 15000
		createdAt   time.Time = time.Now()
		updatedAt   time.Time = time.Now()

		categoryId             uint   = 1
		categoryName           string = "cat 1"
		categorySlug           string = "cat 1"
		categoryAltTitle       string = "cat 1"
		categoryAltDescription string = "cat 1"

		galleryId             uint   = 1
		galleryPath           string = "image-1.jpg"
		galleryAltTitle       string = "image 1"
		galleryAltDescription string = "image 1"
	)

	t.Run("FindError", func(t *testing.T) {
		sqlMock.ExpectQuery("SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1").
			WithArgs(id).WillReturnError(errors.New("can't fetch to mock db"))

		products, err := productRepo.Find(id)

		assert.NotNil(t, err)
		assert.Nil(t, products)
	})

	t.Run("FindSuccess", func(t *testing.T) {
		sqlMock, productRepo := InitConnection()
		sqlMock.ExpectQuery("SELECT * FROM `products` WHERE `products`.`id` = ? AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1").
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "name", "slug", "description", "stock", "price", "created_at", "updated_at"}).
				AddRow(id, name, slug, description, stock, price, createdAt, updatedAt))
		sqlMock.ExpectQuery("SELECT * FROM `product_categories` WHERE `product_categories`.`product_id` = ? AND `product_categories`.`deleted_at` IS NULL").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "category_id", "created_at", "updated_at"}).AddRow(1, id, categoryId, createdAt, updatedAt))
		sqlMock.ExpectQuery("SELECT * FROM `categories` WHERE `categories`.`id` = ? AND `categories`.`deleted_at` IS NULL").WithArgs(categoryId).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "slug", "alt_title", "alt_description", "created_at", "updated_at"}).AddRow(categoryId, categoryName, categorySlug, categoryAltTitle, categoryAltDescription, createdAt, updatedAt))
		sqlMock.ExpectQuery("SELECT * FROM `product_galleries` WHERE `product_galleries`.`product_id` = ? AND `product_galleries`.`deleted_at` IS NULL").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "path", "alt_title", "alt_description", "created_at", "updated_at"}).AddRow(galleryId, id, galleryPath, galleryAltTitle, galleryAltDescription, createdAt, updatedAt))

		product, err := productRepo.Find(id)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.IsType(t, &entity.Product{}, product)
	})
}

func TestProductRepository_FindBySlug(t *testing.T) {
	sqlMock, productRepo := InitConnection()
	var (
		id          uint      = 1
		name        string    = "product 1"
		slug        string    = "product-1"
		description string    = "ini product 1"
		stock       uint      = 22
		price       float64   = 15000
		createdAt   time.Time = time.Now()
		updatedAt   time.Time = time.Now()

		categoryId             uint   = 1
		categoryName           string = "cat 1"
		categorySlug           string = "cat 1"
		categoryAltTitle       string = "cat 1"
		categoryAltDescription string = "cat 1"

		galleryId             uint   = 1
		galleryPath           string = "image-1.jpg"
		galleryAltTitle       string = "image 1"
		galleryAltDescription string = "image 1"
	)

	t.Run("FindBySlugError", func(t *testing.T) {
		sqlMock.ExpectQuery("SELECT * FROM `products` WHERE slug = ? AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1").WithArgs(slug).
			WillReturnError(errors.New("can't fetch to mock db"))

		product, err := productRepo.FindBySlug(slug)

		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("FindBySlugSuccess", func(t *testing.T) {
		sqlMock, productRepo := InitConnection()
		sqlMock.ExpectQuery("SELECT * FROM `products` WHERE slug = ? AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1").WithArgs(slug).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "slug", "description", "stock", "price", "created_at", "updated_at"}).AddRow(id, name, slug, description, stock, price, createdAt, updatedAt))
		sqlMock.ExpectQuery("SELECT * FROM `product_categories` WHERE `product_categories`.`product_id` = ? AND `product_categories`.`deleted_at` IS NULL").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "category_id", "created_at", "updated_at"}).AddRow(1, id, categoryId, createdAt, updatedAt))
		sqlMock.ExpectQuery("SELECT * FROM `categories` WHERE `categories`.`id` = ? AND `categories`.`deleted_at` IS NULL").WithArgs(categoryId).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "slug", "alt_title", "alt_description", "created_at", "updated_at"}).AddRow(categoryId, categoryName, categorySlug, categoryAltTitle, categoryAltDescription, createdAt, updatedAt))
		sqlMock.ExpectQuery("SELECT * FROM `product_galleries` WHERE `product_galleries`.`product_id` = ? AND `product_galleries`.`deleted_at` IS NULL").WithArgs(id).WillReturnRows(sqlmock.NewRows([]string{"id", "product_id", "path", "alt_title", "alt_description", "created_at", "updated_at"}).AddRow(galleryId, id, galleryPath, galleryAltTitle, galleryAltDescription, createdAt, updatedAt))
		product, err := productRepo.FindBySlug(slug)

		assert.Nil(t, err)
		assert.NotNil(t, product)
		assert.IsType(t, &entity.Product{}, product)
	})
}

func TestProductRepository_Insert(t *testing.T) {
	sqlMock, productRepo := InitConnection()
	var (
		id          uint    = 1
		name        string  = "product 1"
		slug        string  = "product-1"
		description string  = "ini product 1"
		stock       uint    = 22
		price       float64 = 15000
		createdAt   AnyTime = AnyTime{}
		updatedAt   AnyTime = AnyTime{}

		//categoryId             uint   = 1
		//categoryName           string = "cat 1"
		//categorySlug           string = "cat 1"
		//categoryAltTitle       string = "cat 1"
		//categoryAltDescription string = "cat 1"
		//
		//galleryId             uint   = 1
		//galleryPath           string = "image-1.jpg"
		//galleryAltTitle       string = "image 1"
		//galleryAltDescription string = "image 1"
	)

	t.Run("InsertError", func(t *testing.T) {
		sqlMock.ExpectExec("INSERT INTO `products` (`name`,`slug`,`description`,`stock`,`price`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?)").WithArgs(name, slug, description, stock, price, createdAt, updatedAt, nil).
			WillReturnError(errors.New("can't fetch to mock db"))

		productEntity := &entity.Product{
			Name:        name,
			Slug:        slug,
			Description: description,
			Stock:       stock,
			Price:       price,
		}

		product, err := productRepo.Insert(productEntity)

		assert.NotNil(t, err)
		assert.Nil(t, product)
	})

	t.Run("InsertSuccess", func(t *testing.T) {
		sqlMock, productRepo := InitConnection()
		sqlMock.ExpectBegin()
		sqlMock.ExpectExec("INSERT INTO `products` (`name`,`description`,`slug`,`stock`,`price`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?,?,?,?)").WithArgs(name, description, slug, stock, price, createdAt, updatedAt, nil).WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectExec("INSERT INTO `product_categories` (`product_id`,`category_id`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?),(?,?,?,?,?) ON DUPLICATE KEY UPDATE `product_id`=VALUES(`product_id`)").WithArgs(id, 1, createdAt, updatedAt, nil, id, 2, createdAt, updatedAt, nil).WillReturnResult(sqlmock.NewResult(1, 1))
		sqlMock.ExpectCommit()
		productEntity := &entity.Product{
			Name:        name,
			Slug:        slug,
			Description: description,
			Stock:       stock,
			Price:       price,
			ProductCategory: []entity.ProductCategory{
				entity.ProductCategory{
					CategoryID: 1,
				},
				entity.ProductCategory{
					CategoryID: 2,
				},
			},
		}
		product, err := productRepo.Insert(productEntity)

		assert.Nil(t, err)
		assert.NotNil(t, product)
	})
}
