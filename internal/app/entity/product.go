package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint           `gorm:"primaryKey"`
	Name       string         `gorm:"type:varchar(255);not null"`
	CategoryID uint           `gorm:"type:int;not null"`
	Category   Category       `gorm:"foreignKey:CategoryID"`
	Price      float64        `gorm:"type:decimal(10,2);not null"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (p *Product) TableName() string {
	return "products"
}
