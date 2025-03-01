package entities

import "time"

// Define field order in struct definition
type DiskonProduk struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID     uint      `gorm:"column:id_produk" json:"id_produk"`
	HargaKonsumen string    `json:"diskon_produk"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func (DiskonProduk) TableName() string {
	return "diskon_produk"
}
