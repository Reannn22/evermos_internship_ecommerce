package entities

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	ID            uint         `json:"id" gorm:"primaryKey;column:id"`
	IDUser        uint         `json:"id_user" gorm:"not null"`
	UserID        uint         `json:"user_id" gorm:"column:id_user"`
	NamaToko      string       `json:"nama_toko" gorm:"column:nama_toko;not null"`
	DeskripsiToko string       `json:"deskripsi_toko" gorm:"column:deskripsi_toko"`
	UrlFoto       string       `json:"url_foto" gorm:"column:url_foto"`
	FotoToko      []StorePhoto `json:"foto_toko" gorm:"foreignKey:IdToko"`
	CreatedAt     *time.Time   `json:"created_at" gorm:"column:created_at"`
	UpdatedAt     *time.Time   `json:"updated_at" gorm:"column:updated_at"`
}

func (Store) TableName() string {
	return "toko"
}
