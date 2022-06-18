package entity

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
	_eProduct "toko/cmd/domain/product/entity"
	_eUser "toko/cmd/domain/user/entity"
)

type Order struct {
	ID            uint `gorm:"primaryKey;autoIncrement;<-:create"`
	UserID        uint
	InvoiceNo     string `gorm:"not null"`
	CustomerName  string `gorm:"not null"`
	CustomerEmail string `gorm:"not null"`
	PaymentMethod string `gorm:"not null"`
	Note          sql.NullString
	Status        uint `gorm:"not null;default:1;comment:'0=Cancel,1=Not Paid,2=Seller Process,3=Shipped,4=Done'"`
	Total         float64

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	User        _eUser.User    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderDetail []*OrderDetail `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type OrderList []*Order

type OrderDetail struct {
	ID                 uint `gorm:"primaryKey;autoIncrement;<-:create"`
	OrderID            uint
	ProductID          uint
	ProductName        string  `gorm:"not null"`
	ProductDescription string  `gorm:"not null;type:text"`
	ProductPrice       float64 `gorm:"not null;default:0"`
	ProductImage       string  `gorm:"-:migration;<-:false"`
	Quantity           uint    `gorm:"not null;default:0"`
	SubTotal           float64 `gorm:"-:migration;<-:false"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Product _eProduct.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Order   Order             `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u *OrderDetail) AfterFind(tx *gorm.DB) (err error) {
	var featuredImage string
	if len(u.Product.ProductGallery) > 0 {
		featuredImage = u.Product.ProductGallery[0].Path
	}
	if u.Quantity > 0 {
		u.SubTotal = float64(u.Quantity) * u.Product.Price
	}
	u.ProductImage = featuredImage
	return
}
