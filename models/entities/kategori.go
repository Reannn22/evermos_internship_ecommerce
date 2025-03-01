package entities

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	NamaCategory string `gorm:"size:255;not null"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (Category) TableName() string {
	return "category"
}
