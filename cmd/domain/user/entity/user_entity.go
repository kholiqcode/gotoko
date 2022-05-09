package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID       uint   `gorm:"primaryKey;autoIncrement;<-:create"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;<-:create"`
	Password string `gorm:"not null"`

	CreatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	DeletedAt gorm.DeletedAt
}

type UserList []*User
