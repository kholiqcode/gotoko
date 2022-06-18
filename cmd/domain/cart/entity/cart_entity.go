package entity

import (
	"gorm.io/gorm"
	"time"
	_eProduct "toko/cmd/domain/product/entity"
	_eUser "toko/cmd/domain/user/entity"
)

type Cart struct {
	ID        uint `gorm:"primaryKey;autoIncrement;<-:create"`
	ProductID uint
	UserID    uint
	Quantity  uint    `gorm:"not null;default:0"`
	SubTotal  float64 `gorm:"-:migration;<-:false"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt

	//Relations
	Product _eProduct.Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User    _eUser.User       `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type CartList []*Cart

func (u *Cart) AfterFind(tx *gorm.DB) (err error) {
	if u.Quantity > 0 {
		u.SubTotal = float64(u.Quantity) * u.Product.Price
	}
	return
}
