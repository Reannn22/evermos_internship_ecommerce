package models

import "time"

type DiskonProdukRequest struct {
	ProductID     uint   `json:"product_id" validate:"required"`
	HargaKonsumen string `json:"harga_konsumen" validate:"required"`
}

type DiskonProdukResponse struct {
	ID            uint      `json:"id"`
	ProductID     uint      `json:"id_produk"`
	HargaKonsumen string    `json:"diskon_produk"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type DiskonProduk struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID     uint      `json:"product_id"`
	HargaKonsumen string    `json:"harga_konsumen"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
