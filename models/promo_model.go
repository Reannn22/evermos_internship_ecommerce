package models

import "time"

type ProductPromoRequest struct {
	IDToko   uint   `json:"id_toko" form:"id_toko"`
	IDProduk uint   `json:"id_produk" form:"id_produk"`
	Promo    string `json:"promo" form:"promo"`
}

type ProductPromoResponse struct {
	ID        uint       `json:"id"`
	IDToko    uint       `json:"id_toko"`
	IDProduk  uint       `json:"id_produk"`
	Promo     string     `json:"promo"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
