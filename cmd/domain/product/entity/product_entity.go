package entity

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement;<-:create"`
	Name        string  `gorm:"not null"`
	Description string  `gorm:"not null;type:text"`
	Slug        string  `gorm:"unique;not null;size:50"`
	Stock       uint    `gorm:"not null;default:0"`
	Price       float64 `gorm:"not null;default:0"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	ProductGallery  []ProductGallery  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductCategory []ProductCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ProductList []*Product

type ProductGallery struct {
	ID             uint `gorm:"primaryKey;autoIncrement;<-:create"`
	ProductID      uint
	Path           string `gorm:"not null"`
	AltTitle       sql.NullString
	AltDescription sql.NullString

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//	Relations
	Product Product `gorm:"foreignKey:ProductID;references:ID"`
}

type Category struct {
	ID             uint   `gorm:"primaryKey;autoIncrement;<-:create"`
	Name           string `gorm:"not null"`
	Slug           string `gorm:"unique;not null;size:50"`
	AltTitle       sql.NullString
	AltDescription sql.NullString

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//	Relations
	ProductCategory []ProductCategory `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

// Pivot Entity
type ProductCategory struct {
	ID         uint `gorm:"primaryKey;autoIncrement;<-:create"`
	ProductID  uint
	CategoryID uint

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Product  Product  `gorm:"foreignKey:ProductID;references:ID"`
	Category Category `gorm:"foreignKey:CategoryID;references:ID"`
}
