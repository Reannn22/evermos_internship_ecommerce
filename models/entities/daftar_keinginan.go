package entities

import "time"

type Wishlist struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	IDToko    uint       `json:"id_toko" gorm:"column:id_toko"`
	IDProduk  uint       `json:"id_produk" gorm:"column:id_produk"`
	Store     Store      `json:"toko" gorm:"foreignKey:IDToko"`
	Product   Product    `json:"produk" gorm:"foreignKey:IDProduk"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func (Wishlist) TableName() string {
	return "daftar_keinginan_belanja"
}
