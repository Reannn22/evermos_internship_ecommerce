package entities

import (
	"time"
)

type StorePhoto struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	IdFoto    uint       `json:"-" gorm:"column:id_foto"` // Add json:"-" to hide from response
	IdToko    uint       `json:"-" gorm:"column:id_toko"` // Add json:"-" to hide from response
	URL       string     `json:"url"`
	Photo     string     `json:"photo"`
	FileName  string     `json:"file_name"`
	FileType  string     `json:"file_type"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (StorePhoto) TableName() string {
	return "foto_toko"
}

