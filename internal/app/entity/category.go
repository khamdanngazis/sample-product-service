package entity

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"-"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (c *Category) TableName() string {
	return "categories"
}
